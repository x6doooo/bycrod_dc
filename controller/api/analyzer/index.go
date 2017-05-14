package analyzer

import (
    "gopkg.in/gin-gonic/gin.v1"
    "bycrod_dc/service/analyzer"
    "net/http"
    "bycrod_dc/service/util"
)

func Signals(ctx *gin.Context) {
    data := analyzer.Signal()
    ctx.JSON(http.StatusOK, util.OkResponse(data))
}