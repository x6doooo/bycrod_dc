package analyzer

import (
    "github.com/x6doooo/talib"
    "bycrod_dc/service/util"
    "bycrod_dc/service/stock"
)

func Signal() interface{} {
    //mongo.DB.CollectionNames()
    return dailyDataHandler("AMD")
}

func dailyDataHandler(code string) interface{} {
    data := stock.QueryTimeSeriesDataByDate(code, "daily", "", "")

    //signals, outBegIdx, outNBElement := talib.Cdl2Crows(data.Open, data.High, data.Low, data.Close)
    //signals, outBegIdx, outNBElement := talib.Cdl3BlackCrows(data.Open, data.High, data.Low, data.Close)
    //signals, outBegIdx, outNBElement := talib.Cdl3Inside(data.Open, data.High, data.Low, data.Close)
    //signals, outBegIdx, outNBElement := talib.Cdl3LineStrike(data.Open, data.High, data.Low, data.Close)
    signals, outBegIdx, outNBElement := talib.Cdl3Outside(data.Open, data.High, data.Low, data.Close)
    signals = talib.FormatInt32(signals, outBegIdx, outNBElement)

    util.Logger.Info("%v", signals)
    return signals
}
