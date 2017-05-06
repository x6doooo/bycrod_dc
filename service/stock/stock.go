package stock


import (
    "bycrod_dc/conf"
    "github.com/x6doooo/snowball"
    "gopkg.in/mgo.v2"
    "bycrod_dc/service/common/mongo"
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








