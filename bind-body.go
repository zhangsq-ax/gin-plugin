package gin_plugin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func BindBody(ctx *gin.Context, obj interface{}) error {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, obj)
}