package gin_plugin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ResponseJSON struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponseJSON() *ResponseJSON {
	return &ResponseJSON{
		Status:  200,
		Message: "ok",
	}
}

func (res *ResponseJSON) Marshal() []byte {
	data, _ := json.Marshal(res)
	return data
}

func (res *ResponseJSON) SetStatus(status int) *ResponseJSON {
	res.Status = status
	return res
}

func (res *ResponseJSON) SetMessage(message string) *ResponseJSON {
	if message != "" {
		res.Message = message
	}
	return res
}

func (res *ResponseJSON) SetData(data interface{}) *ResponseJSON {
	res.Data = data
	return res
}

func (res *ResponseJSON) Set(status int, message string, data ...interface{}) *ResponseJSON {
	res.SetStatus(status)
	res.SetMessage(message)
	if len(data) > 0 {
		res.SetData(data[0])
	}
	return res
}

func (res *ResponseJSON) End(ctx *gin.Context, code int) {
	ctx.JSON(code, res)
}
