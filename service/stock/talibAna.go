package stock

import (
    "github.com/x6doooo/talib"
)


func TalibDispatcher(functionName, code, dataType, startDate, endDate string) interface{} {
    res := QueryTimeSeriesDataByDate(code, dataType, startDate, endDate)
    switch(functionName) {
    case "Cdl2Crows":
        data := talib.Cdl2Crows(res.Open, res.High, res.Low, res.Close)
        return data
    case "Cdl3BlackCrows":
        data := talib.Cdl3BlackCrows(res.Open, res.High, res.Low, res.Close)
        return data
    case "Cdl3Inside":
        data := talib.Cdl3Inside(res.Open, res.High, res.Low, res.Close)
        return data
    }
    return nil
}


