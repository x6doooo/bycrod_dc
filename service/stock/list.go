package stock

import (
    "bycrod_dc/service/util"
    "bycrod_dc/service/common/task"
    "math"
    "strings"
    "bycrod_dc/model/stock"
    "github.com/x6doooo/gout"
    "gopkg.in/mgo.v2/bson"
    "bycrod_dc/service/common/mongo"
)


func workerForGetDetail(jobs <- chan string, results chan<-[]map[string]interface{}) {
    for codes := range jobs {
        //util.Logger.Info("getDetailByCodes: %s", codes)
        details := SbClient.GetDetail(codes)
        results <- details
    }
}

func getDetails(codes []string) (details []stock.Record) {
    workerNum := 2
    jobNum := int(math.Ceil(float64(len(codes)) / 50.0))

    jobs := make(chan string, jobNum)
    results := make(chan []map[string]interface{}, jobNum)
    defer close(jobs)
    defer close(results)
    for idx := 0; idx < workerNum; idx++ {
        go workerForGetDetail(jobs, results)
    }

    for idx := 0; idx < jobNum; idx++ {
        startIdx := idx * 50
        endIdx := (idx + 1) * 50
        if endIdx > len(codes) {
            endIdx = len(codes)
        }
        util.Logger.Info("startIdx & endIdx => %d:%d", startIdx, endIdx)
        codeSlice := codes[startIdx:endIdx]
        jobs <- strings.Join(codeSlice, ",")
    }

    for idx := 0; idx < jobNum; idx++ {
        res := <-results
        util.Logger.Info("received results: %d", len(res))
        for _, r := range res {
            record := stock.Record{}
            err := gout.Map2Struct(r, &record, false)
            if err == nil {
                details = append(details, record)
            } else {
                util.Logger.Error("%v", err);
            }
        }
    }
    return
}


func LoadRecordListFromDb(condition bson.M, sort string, skip int, limit int) (list []stock.Record) {

    q := StockListCollection.Find(condition)

    if sort != "" {
        q = q.Sort(sort)
    }
    if skip != 0 {
        q = q.Skip(skip)
    }
    if limit != 0 {
        q = q.Limit(limit)
    }

    q.All(&list)
    return
}

func LoadRecordsFromDb(condition bson.M, sort string, skip int, limit int) (recordMap map[string]stock.Record) {
    list := LoadRecordListFromDb(condition, sort, skip, limit)
    recordMap = make(map[string]stock.Record)
    for _, item := range list {
        recordMap[item.Code] = item;
    }
    return
}

func Count(condition bson.M) (int, error) {
    return StockListCollection.Find(condition).Count()
}

func merge(recordMap map[string]stock.Record, details []stock.Record) (results []interface{}) {

    for _, detailCast := range details {

        if detailCast.Code == "" {
            util.Logger.Info("detailCast has no Code: %v", detailCast)
            continue
        }

        if record, ok := recordMap[detailCast.Code]; ok {
            detailCast.Watching = record.Watching
            detailCast.Id = record.Id
            delete(recordMap, detailCast.Code)
        } else {
            detailCast.Watching = false
        }
        results = append(results, detailCast)
    }

    for _, record := range recordMap {
        record.NoUpdate = true
        results = append(results, record)
    }

    util.Logger.Info("merge result size: %d", len(results))

    return results

}

/*
    更新整个列表
 */
func UpdateList() error {

    taskName := "UpdateStockList"

    err := task.Lock(taskName)
    if err != nil {
        return err
    }

    defer task.Unlock(taskName)
    codes := SbClient.GetCodeList()
    util.Logger.Info("getcodes from xueqiu: %d", len(codes))

    // test
    //codes = codes[0:155]
    util.Logger.Info("codes size: %d", len(codes))

    details := getDetails(codes)
    util.Logger.Info("details size: %d", len(details))

    recordMap := LoadRecordsFromDb(bson.M{}, "", 0, 0)
    list := merge(recordMap, details)
    StockListCollection.RemoveAll(bson.M{})
    StockListCollection.Insert(list...)
    return nil
}


func GetWatchingCodes() (codes []string) {
    list := []map[string]string{}
    StockListCollection.Find(bson.M{
        "watching": true,
    }).Select(bson.M{
        "code": 1,
    }).All(&list)
    for _, item := range list {
        codes = append(codes, item["code"])
    }
    return
}

func CleanUnwatchingCodeDataCollection() {

    taskName := "CleanCodeCollection"
    task.Lock(taskName)
    defer task.Unlock(taskName)

    names, err := mongo.DB.CollectionNames()
    if err != nil {
        panic(err)
    }

    watchingCodes := GetWatchingCodes()
    watchingCodeMap := make(map[string]bool)
    for _, wcName := range watchingCodes {
        watchingCodeMap[wcName] = true
    }

    for _, name := range names {
        util.Logger.Info("name: %s", name)
        idx := strings.Index(name, "code_")
        if idx != 0 {
            continue
        }
        codeName := strings.Split(name, "_")[1]
        util.Logger.Info("codeName: %s", codeName)
        if _, ok := watchingCodeMap[codeName]; ok {
            continue
        }
        util.Logger.Info("remove: %s", name)
        mongo.DB.C(name).DropCollection()
    }
}


