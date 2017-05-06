package yahoo

import (
    "bycrod_dc/service/stock"
    "time"
    "bycrod_dc/service/util"
    "bycrod_dc/service/common/task"
    "sync"
    "gopkg.in/mgo.v2/bson"
    "bycrod_dc/service/common/mongo"
    "errors"
)

const (
    interval_daily = "1d"
    the_range_daily = "3072d"

    interval_hourly = "1h"
    the_range_hourly = "1024d"

    interval_minutely = "1m"
    the_range_minutely = "100d"

    interval_realtime = "1m"
    the_range_realtime = "5m"
)

type ByHigh []QuoteItem
func (me ByHigh) Len() int {
    return len(me)
}
func (me ByHigh) Swap(i, j int) {
    me[i], me[j] = me[j], me[i]
}
func (me ByHigh) Less(i, j int) bool {
    return me[i].High < me[i].High
}

type ByLow []QuoteItem
func (me ByLow) Len() int {
    return len(me)
}
func (me ByLow) Swap(i, j int) {
    me[i], me[j] = me[j], me[i]
}
func (me ByLow) Less(i, j int) bool {
    return me[i].Low < me[i].Low
}


func Load(dataType string) (err error) {

    taskName := "loadStockDataFromYahooApi"
    err = task.Lock(taskName)
    defer task.Unlock(taskName)
    if err != nil {
        util.Logger.Error("%v", err)
        return
    }

    codes := stock.GetWatchingCodes()
    util.Logger.Info("watching codes size: %d", len(codes))
    startTime := time.Now()
    count := 0
    for {
        codes = loop(codes, startTime, dataType)
        util.Logger.Info("failed size: %d", len(codes))
        if len(codes) == 0 {
            break
        }
        count += 1
        if count > 10 {
            break
        }
    }

    util.Logger.Info("yahoo data fetch done!")
    return
}

func loop(codes []string, startTime time.Time, dataType string) (failedCodes []string) {
    var interval string
    var the_range string
    switch dataType {
    case "daily":
        interval = interval_daily
        the_range = the_range_daily
    case "hourly":
        interval = interval_hourly
        the_range = the_range_hourly
    case "minutely":
        interval = interval_minutely
        the_range = the_range_minutely
    case "realtime":
        interval = interval_realtime
        the_range = the_range_realtime
    default:
        interval = interval_daily
        the_range = the_range_daily
    }
    processNum := 2
    step := len(codes) / processNum

    wg := sync.WaitGroup{}
    wg.Add(processNum)
    for i := 0; i < processNum; i++ {
        var codesOfStep []string
        if i == processNum - 1 {
            codesOfStep = codes[i * step : ]
        } else {
            codesOfStep = codes[i * step : (i + 1) * step]
        }
        go fetchRoutine(codesOfStep, interval, the_range, &wg, &failedCodes, dataType, startTime)
    }

    wg.Wait()


    return
}


func fetchRoutine(codes []string, interval string, the_range string,
wg *sync.WaitGroup,
failedCodes *([]string), dataType string, startTime time.Time,) {

    defer wg.Done()
    size := len(codes);
    for idx, code := range codes {

        collectionName := util.CollectionName(code, dataType)
        //collectionName := "code_" + code + "_" + dataType

        util.Logger.Info("%s %d/%d %s", code, idx, size, time.Since(startTime).String())

        startTimeOfCurrentCode := time.Now()
        results, err := Get(code, interval, the_range)
        timeUsedOfFetch := time.Now()
        if err != nil {
            util.Logger.Info("%s failed: %s", code, err.Error())
            *failedCodes = append(*failedCodes, code)
            continue
        }
        timeUsedOfCompute := time.Now()

        var theLastTs int64 = 0

        var dataListHasBeenInserted []QuoteItem
        mongo.DB.C(collectionName).Find(bson.M{}).Sort("-ts").All(&dataListHasBeenInserted)
        if len(dataListHasBeenInserted) != 0 {
            theLastTs = dataListHasBeenInserted[0].Ts
        }

        dataList, err := handle(results, theLastTs)
        if err != nil {
            continue
        }
        if len(dataList) > 0 {
            mongo.DB.C(collectionName).Insert(dataList...)
        }
        timeUsedOfInsert := time.Now()

        util.Logger.Info(" - fetch: %s", timeUsedOfFetch.Sub(startTimeOfCurrentCode).String())
        util.Logger.Info(" - compute: %s", timeUsedOfCompute.Sub(timeUsedOfFetch).String())
        util.Logger.Info(" - insert: %s", timeUsedOfInsert.Sub(timeUsedOfCompute).String())
    }

}



func InitBaseValue(idx int, ts int64, quotes QuoteData) (item QuoteItem, err error) {

    item.Ts = ts
    d := time.Unix(ts, 0)
    //item.DateTime = d.UTC().Format("2006-01-02 15:04:05")
    item.Date = d.UTC().Format("2006-01-02")

    theVolume := quotes.Volume[idx]
    if theVolume != nil {
        item.Volume = *theVolume
    } else {
        err = errors.New("volume error")
        return
    }
    theOpen := quotes.Open[idx]
    if theOpen != nil {
        item.Open = *theOpen
    } else {
        err = errors.New("open error")
        return
    }
    theClose := quotes.Close[idx]
    if theClose != nil {
        item.Close = *theClose
    } else {
        err = errors.New("close error")
        return
    }
    theHigh := quotes.High[idx]
    if theHigh != nil {
        item.High = *theHigh
    } else {
        err = errors.New("high error")
        return
    }
    theLow := quotes.Low[idx]
    if theLow != nil {
        item.Low = *theLow
    } else {
        err = errors.New("low error")
        return
    }
    return
}

func handle(respData RespResult, theLastTs int64) (listInterface []interface{}, err error) {

    resultArr := respData.Chart.Result
    if len(resultArr) == 0 {
        err = errors.New("results is empty")
        return
    }

    result := respData.Chart.Result[0]
    timestamps := result.Timestamp
    quotes := result.Indicators["quote"][0]

    // init
    list := make([]QuoteItem, 0, len(timestamps))
    for idx, ts := range timestamps {

        if ts <= theLastTs {
            continue
        }
        item, err := InitBaseValue(idx, ts, quotes)
        if err != nil {
            continue
        }

        list = append(list, item)

        listInterface = append(listInterface, item)

    }

    return
}
