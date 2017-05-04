package util

import "gopkg.in/gin-gonic/gin.v1"

func ErrorResponse(data interface{}) gin.H {
    return gin.H{
        "code": 1,
        "data": data,
    }
}

func OkResponse(data interface{}) gin.H {
    return gin.H{
        "code": 0,
        "data": data,
    }
}
