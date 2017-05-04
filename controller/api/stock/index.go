package stock_list

import (
    "gopkg.in/gin-gonic/gin.v1"
    stockService "bycrod_dc/service/stock"
    "bycrod_dc/service/util"
    "net/http"
    "gopkg.in/mgo.v2/bson"
    "strconv"
)


func GetStockList(ctx *gin.Context) {

    code := ctx.DefaultQuery("code", "")
    exchange := ctx.DefaultQuery("exchange", "")
    page := ctx.DefaultQuery("page", "1")
    rows := ctx.DefaultQuery("rows", "50")
    order := ctx.DefaultQuery("order", "asc")
    orderBy := ctx.DefaultQuery("orderBy", "")

    sort := ""
    if orderBy != "" {
        if order == "asc" {
            sort = orderBy
        }
        if order == "desc" {
            sort = "-" + orderBy
        }
    }

    condition := bson.M{}
    if code != "" {
        condition["code"] = code
    }
    if exchange != "" {
        condition["exchange"] = exchange
    }

    pageInt, _ := strconv.Atoi(page)
    rowsInt, _ := strconv.Atoi(rows)
    skip := (pageInt - 1) * rowsInt
    if skip < 0 {
        skip = 0
    }

    count, _ := stockService.Count(condition)

    list := stockService.LoadRecordListFromDb(condition, sort, skip, rowsInt)

    result := gin.H{
        "count": count,
        "page": pageInt,
        "rows": rowsInt,
        "data": list,
    }
    ctx.JSON(http.StatusOK, util.OkResponse(result))
}

func UpdateList(ctx *gin.Context) {
    stockService.UpdateList()
    ctx.JSON(200, gin.H{
        "code": 0,
    })
}

func ModifyWatchingState(ctx *gin.Context) {
    code := ctx.Query("code")
    state := ctx.Query("state")

    stateBool := false
    if state == "true" {
        stateBool = true
    }
    err := stockService.UpdateWatchingState(code, stateBool)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError,
            util.ErrorResponse(err.Error()))
        return
    }
    ctx.JSON(http.StatusOK, util.OkResponse("ok"))
}

