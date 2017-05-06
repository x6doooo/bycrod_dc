package stock

import (
    "github.com/x6doooo/talib"
)


func TalibDispatcher(functionName, code, dataType, startDate, endDate string) interface{} {
    res := QueryTimeSeriesDataByDate(code, dataType, startDate, endDate)
    switch(functionName) {
    case "Cdl2Crows":
        data, outBegIdx, outNBElement := talib.Cdl2Crows(res.Open, res.High, res.Low, res.Close)
        data = talib.FormatInt32(data, outBegIdx, outNBElement)
        return data
    case "Cdl3BlackCrows":
        data, outBegIdx, outNBElement := talib.Cdl3BlackCrows(res.Open, res.High, res.Low, res.Close)
        data = talib.FormatInt32(data, outBegIdx, outNBElement)
        return data
    }
    return nil
}


