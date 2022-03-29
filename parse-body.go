package gin_plugin

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangsq-ax/logs"
	"go.uber.org/zap"
)

var logger = logs.NewJSONLogger("gin-plugin")

func ParseBody(ctx *gin.Context, res *ResponseJSON, req interface{}) (requestBody interface{}, hasError bool) {
	hasError = false
	//err := ctx.BindJSON(req)
	//err := ctx.Bind(req)
	err := BindBody(ctx, req)
	if err != nil {
		logger.Warnw("invalid request", zap.Error(err))
		res.Set(400, "invalid request")
		hasError = true
		return
	}
	requestBody = req
	return
}
