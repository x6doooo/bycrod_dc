package strategy

import (
    "bycrod_dc/service/util"
    "bycrod_dc/service/common/mongo"
    "gopkg.in/mgo.v2/bson"
    "bycrod_dc/model/stock"
)

func LoadData(code, dataType string, points int) (res []stock.Record, err error) {
    collectionName := util.CollectionName(code, dataType)
    err = mongo.DB.C(collectionName).Find(bson.M{
    }).Sort(bson.M{
        "ts": -1,
    }).Limit(points).All(&res)
}

func FormatData(data []stock.Record, fields []string) (res map[string][]float64) {
    res = make(map[string][]float64)

    // init
    size := len(data)
    for _, f := range fields {
        res[f] = make([]float64, size)
    }

    for idx, item := range data {
        for _, f := range fields {
            res[f][idx] = item[f]
        }
    }

}