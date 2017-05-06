package yahoo

import (
    "gopkg.in/mgo.v2/bson"
    "reflect"
)

type QuoteItem struct {
    Id                     bson.ObjectId `bson:"_id,omitempty"`
    Ts                     int64
    Date                   string
    Volume                 uint64
    High                   float64
    Low                    float64
    Open                   float64
    Close                  float64
}

func (me *QuoteItem) SetFloat64ByFieldName(field string, value float64) (ok bool) {
    r := reflect.ValueOf(me)
    f := r.Elem().FieldByName(field)
    ok = f.IsValid() && f.CanSet() && f.Kind() == reflect.Float64
    if ok {
        f.SetFloat(value)
    }
    return
}

func (me *QuoteItem) GetFloat64ByFieldName(field string) (value float64, ok bool) {
    r := reflect.ValueOf(me)
    f := r.Elem().FieldByName(field)
    ok = f.IsValid()
    if ok {
        value = f.Float()
    }
    return
}
type QuoteData struct {
    Volume []*uint64
    High   []*float64
    Low    []*float64
    Open   []*float64
    Close  []*float64
}

type ResultItem struct {
    Meta       interface{}
    Timestamp  []int64
    Indicators map[string]([]QuoteData)
}

type ChartItem struct {
    Result []ResultItem
    Error  *string
}

type RespResult struct {
    Chart ChartItem
}
