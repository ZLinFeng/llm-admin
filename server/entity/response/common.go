package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS  = 10_000
	AUTH_ERR = 10_001
	COM_ERR  = 10_002
)

const (
	SUCC_MSG = "Success"
	Fail_MSG = "Fail"
)

type Response struct {
	Code int         `json:code`
	Msg  string      `json:msg`
	Data interface{} `json:data`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	//Result()
}
