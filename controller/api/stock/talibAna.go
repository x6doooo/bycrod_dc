package stock_list

import (
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
    "errors"
    "bycrod_dc/service/util"
    stockService "bycrod_dc/service/stock"
)

func TalibHandler(ctx *gin.Context) {
    code := ctx.DefaultQuery("code", "")
    dataType := ctx.DefaultQuery("dataType", "daily")
    startDate := ctx.DefaultQuery("startDate", "")
    endDate := ctx.DefaultQuery("endDate", "")
    functionName := ctx.Param("function")

    if code == "" {
        panic(errors.New("need stock code"))
        return;
    }

    res := stockService.TalibDispatcher(functionName, code, dataType, startDate, endDate)
    ctx.JSON(http.StatusOK, util.OkResponse(res))
}



