package server

import (
    "time"
    "gopkg.in/gin-gonic/gin.v1"
    "github.com/x6doooo/err_handler"
    "bycrod_dc/service/util"
    "bycrod_dc/conf"
    "net/http"
    stockCtrl "bycrod_dc/controller/api/stock"
    commonCtrl "bycrod_dc/controller/api/common"
)


func RequestLog() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        ip := c.ClientIP()

        c.Next()
        end := time.Now()
        latency := end.Sub(start)
        util.Logger.Info("[%d] %s %s %s %s",
            c.Writer.Status(), ip,  c.Request.Method, c.Request.RequestURI, latency.String())
    }
}
//
func ErrHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        var err error
        defer err_handler.Recover(&err, func() {
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "code": -1,
                    "data": err.Error(),
                })
                c.Abort()
            }
        })
        c.Next()
    }
}


func Start() {

    websocketServer := newWebsocketServer()


    engine := gin.New()
    engine.Use(ErrHandler())
    engine.Use(gin.Recovery())

    engine.GET("/socket/", gin.WrapH(websocketServer))

    // request log
    engine.Use(RequestLog())

    apiRouter := engine.Group("/api")
    {
        // 列表
        apiRouter.GET("/stock/list", stockCtrl.GetStockList)
        // 更新列表
        apiRouter.GET("/stock/list/update", stockCtrl.UpdateList)
        // 更新watching状态
        apiRouter.PUT("/stock/watching", stockCtrl.ModifyWatchingState)
        // 删除没有关注的时序collection
        apiRouter.DELETE("/stock/unwatching", stockCtrl.CleanUnwatchingCodeDataCollection)
        // 抓取数据
        apiRouter.GET("/stock/timeSeriesData/load", stockCtrl.LoadStockData)

        //--- talib ---
        apiRouter.GET("/stock/talib/:function", stockCtrl.TalibHandler)


        // 查询时序数据
        apiRouter.GET("/stock/timeSeriesData/query", stockCtrl.QueryTimeSeriesData)
        // 获取task的状态
        apiRouter.GET("/common/task/status", commonCtrl.GetTaskStatus)
    }

    util.Logger.Info("server start! %s", conf.MainConf.Server.Addr)
    engine.Run(conf.MainConf.Server.Addr)
}

