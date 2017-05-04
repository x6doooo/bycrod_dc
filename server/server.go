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
                    "data": err.Error(),
                })
                c.Abort()
            }
        })
        c.Next()
    }
}


func Start() {
    engine := gin.New()
    engine.Use(ErrHandler())
    engine.Use(gin.Recovery())

    // request log
    engine.Use(RequestLog())

    apiRouter := engine.Group("/api")
    {
        apiRouter.GET("/stock/list", stockCtrl.GetStockList)
        apiRouter.GET("/stock/list/update", stockCtrl.UpdateList)
        apiRouter.PUT("/stock/watching", stockCtrl.ModifyWatchingState)

        apiRouter.GET("/common/task/status", commonCtrl.GetTaskStatus)
    }

    util.Logger.Info("server start! %s", conf.MainConf.Server.Addr)
    engine.Run(conf.MainConf.Server.Addr)
}

