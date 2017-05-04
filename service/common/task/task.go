/**
    不能并发执行的任务，通过mongo标记状态;
 */
package task

import (
    "bycrod_dc/service/common/mongo"
    "gopkg.in/mgo.v2/bson"
    "gopkg.in/mgo.v2"
)

type taskModel struct {
    Id    bson.ObjectId `bson:"_id,omitempty"`
    Key   string
    Value bool
}

const (
    collectionName = "tasks"
)

var (
    collectionCli *mgo.Collection
)

func init() {
    collectionCli = mongo.DB.C(collectionName)

    var tasks []taskModel
    collectionCli.Find(bson.M{}).All(&tasks)
    for _, t := range tasks {
        collectionCli.Update(bson.M{
            "_id": t.Id,
        }, bson.M{
            "$set": bson.M{
                "value": false,
            },
        })
    }
}

func Lock(key string) error {
    _, err := collectionCli.Upsert(bson.M{
        "key": key,
        "value": false,
    }, bson.M{
        "$set": bson.M{
            "value": true,
        },
    })
    return err
}

func Unlock(key string) error {
    err := collectionCli.Update(bson.M{
        "key": key,
    }, bson.M{
        "$set": bson.M{
            "value": false,
        },
    })
    return err
}

func Status() (map[string]interface{}, error) {
    list := []bson.M{}
    status := map[string]interface{}{}
    err := collectionCli.Find(bson.M{}).All(&list)
    if err != nil {
        return status, err
    }
    for _, item := range list {
        k, ok1 := item["key"]
        v, ok2 := item["value"]
        if ok1 && ok2 {
            status[k.(string)] = v
        }
    }
    return status, nil
}


