package stock

import "gopkg.in/mgo.v2/bson"


func updateRecord(condition bson.M, data bson.M) error {
    return StockListCollection.Update(condition, data)
}


func UpdateWatchingState(code string, state bool) error {
    return updateRecord(bson.M{
        "code": code,
    }, bson.M{
        "$set": bson.M{
            "watching": state,
        },
    })
}
