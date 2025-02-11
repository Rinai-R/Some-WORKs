package Middleware

import (
	"context"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Sentinel(c context.Context, ctx *app.RequestContext) {
	a, err := sentinel.Entry("kitex-etcd")
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请求太多",
		})
		ctx.Abort()
		return
	}
	defer a.Exit()
	ctx.Next(c)

}
