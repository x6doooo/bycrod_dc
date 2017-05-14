package strategy

import "reflect"

type Strategy interface {
    Init()
    HandleBar()
}


var (
    registry = make(map[string]reflect.Type)
)


func init() {
    registry["001_ADX_DMI"] = reflect.TypeOf(ADX_DMI{})
}

func makeStrategyInstance(name string) Strategy {
    v := reflect.New(registry[name]).Elem()
    return v.Interface().(Strategy)
}


