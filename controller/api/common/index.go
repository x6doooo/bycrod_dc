package common

import (
    "gopkg.in/gin-gonic/gin.v1"
    "bycrod_dc/service/common/task"
    "net/http"
    "bycrod_dc/service/util"
)

func GetTaskStatus(ctx *gin.Context) {
    status, err := task.Status()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError,
            util.ErrorResponse(err.Error()))
        return
    }
    ctx.JSON(http.StatusOK,
        util.OkResponse(status))
}