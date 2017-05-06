package stock

import (
    "bycrod_dc/service/common/mongo"
    "bycrod_dc/service/util"
    "gopkg.in/mgo.v2/bson"
)




type QueryDataResult struct {
    Open   []float64
    Close  []float64
    High   []float64
    Low    []float64
    Volume []float64
    Ts     []int64
    Date   []string
}

var (
    baseConditionSort = bson.M{
        "$sort": bson.M{
            "ts": 1,
        },
    }
    baseConditionFields = bson.M{
        "$group": bson.M{
            "_id": nil,
            "open": bson.M{
                "$push": "$open",
            },
            "close": bson.M{
                "$push": "$close",
            },
            "high": bson.M{
                "$push": "$high",
            },
            "low": bson.M{
                "$push": "$low",
            },
            "volume": bson.M{
                "$push": "$volume",
            },
            "ts": bson.M{
                "$push": "$ts",
            },
            "date": bson.M{
                "$push": "$date",
            },
        },
    }
)

func QueryDailyData(collectionName string) QueryDataResult {
    condition := []bson.M{
        //bson.M{
        //    "$match": bson.M{
        //        "date": bson.M{
        //            "$lt": dateStr,
        //        },
        //    },
        //},
        baseConditionSort,
        baseConditionFields,
    }
    res := QueryDataResult{}
    mongo.DB.C(collectionName).Pipe(condition).One(&res)
    return res
}


func QueryTimeSeriesDataByDate(code, dataType, startDate, endDate string) (res QueryDataResult) {
    collectionName := util.CollectionName(code, dataType)
    conditions := []bson.M{}
    if (startDate != "" && endDate != "") {
        dateRangeCondition := bson.M{
            "$match": bson.M{
                "date": bson.M{
                    "$gte": startDate,
                    "$lte": endDate,
                },
            },
        }
        conditions = append(conditions, dateRangeCondition)
    }
    conditions = append(conditions, baseConditionSort);
    conditions = append(conditions, baseConditionFields);
    util.Logger.Info("query condition: %v", conditions)

    mongo.DB.C(collectionName).Pipe(conditions).One(&res)
    return
}
