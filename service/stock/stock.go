package stock


import (
    "bycrod_dc/conf"
    "github.com/x6doooo/snowball"
    "gopkg.in/mgo.v2"
    "bycrod_dc/service/common/mongo"
    "gopkg.in/mgo.v2/bson"
    "bycrod_dc/service/util"
)


const (
    stockListCollectionName = "stock_list"
)


var (
    SbClient *snowball.Client
    StockListCollection *mgo.Collection
)


func init() {

    StockListCollection = mongo.DB.C(stockListCollectionName)

    sbConfig := conf.MainConf.Xueqiu
    SbClient = snowball.New(sbConfig.Username, sbConfig.Password)
    SbClient.Login()

}

func DropUnwatchingTsData() (error) {

    list := []map[string]string{}
    StockListCollection.Find(bson.M{
        "watching": false,
    }).Select(bson.M{
        "code": 1,
    }).All(&list)
    util.Logger.Info("/stock/timeSeriesData/unwatching,  size: %d", len(list))


    needRemove := make(map[string]bool)
    names, err := mongo.DB.CollectionNames()

    if err != nil {
        return err
    }

    for _, name := range names {
        needRemove[name] = true
    }

    for _, item := range list {
        collectionName := util.CollectionName(item["code"], "daily")
        if _, ok := needRemove[collectionName]; ok {
            util.Logger.Info("drop %s", collectionName)
            mongo.DB.C(collectionName).DropCollection()
        }
    }

    return nil

}







