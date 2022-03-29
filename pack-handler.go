package gin_plugin

import "github.com/gin-gonic/gin"

func PackHandler(handler func(ctx *gin.Context, res *ResponseJSON)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := NewResponseJSON()
		defer res.End(ctx, 200)
		handler(ctx, res)
	}
}
