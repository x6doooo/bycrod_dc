package strategy

import (
    stockService "bycrod_dc/service/stock"
    "talib"
)

//原文地址
//https://www.ricequant.com/community/topic/273/#share-source-code_content_1370_847762

// 仓位100为满  实际就是100股

type ADX_DMI struct{

    Code string // 执行策略的股票code
    Benchmark string // 用于对比的股票code
    ADXPERIOD int
    ShortMA int
    LongMa int
}

func (me ADX_DMI) Init() {
    me.Code = "AMD"
    me.ADXPERIOD = 14
    me.ShortMA = 5
    me.LongMa = 20
}

func (me ADX_DMI) HandleBar(ctx *Context) {

    data := stockService.QueryDailyData(me.Code, "daily", 100)

    adx := talib.Adx(data.High, data.Low, data.Close, me.ADXPERIOD)
    //adxr := talib.Adxr(data.High, data.Low, data.Close, me.ADXPERIOD)
    plusDi := talib.PlusDi(data.High, data.Low, data.Close, me.ADXPERIOD)
    minusDi := talib.MinusDi(data.High, data.Low, data.Close, me.ADXPERIOD)
    shortMa := talib.Sma(data.Close, me.ShortMA)
    longMa := talib.Sma(data.Close, me.LongMa)

    curentPrice := data.Close[len(data.Close) - 1]

    if (shortMa[len(shortMa) - 1] > longMa[len(longMa) - 1]) &&
        (shortMa[len(shortMa) - 2] < longMa[len(longMa) - 2]) &&
        (adx[len(adx)-1] > adx[len(adx)-2]) &&
        (plusDi[len(plusDi - 1)] > minusDi[len(minusDi) - 1]) {
        ctx.Order(me.Code, curentPrice, 1.0)
    }

    if (shortMa[len(shortMa) - 1] < longMa[len(longMa) - 1]) &&
        (shortMa[len(shortMa) - 2] > longMa[len(longMa) - 2]) &&
        (adx[len(adx)-1] < adx[len(adx)-2]) &&
        (plusDi[len(plusDi - 1)] < minusDi[len(minusDi) - 1]) {
        ctx.Order(me.Code, curentPrice, 0)
    }



}


//func (me ADX_DMI) handleOneCode(code string) {
//
//}