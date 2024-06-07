package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	SUCCEES = 0
	ERROR   = -1
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
func Succees(data any, msg string, c *gin.Context) {
	Result(SUCCEES, data, msg, c)
}
func SucceesWithData(data any, c *gin.Context) {
	Result(SUCCEES, data, "成功", c)
}
func SucceesMessage(msg string, c *gin.Context) {
	Result(SUCCEES, map[string]any{}, msg, c)
}
func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailMessage(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}

func FailCode(code int, c *gin.Context) {
	msg, ok := ErrorMaps[ErrorCode(code)]
	if ok {
		Result(code, map[string]any{}, msg, c)
		return
	}
	Result(ERROR, map[string]any{}, "神金,害我笑了,发生了蜜汁错误", c)
}
