package MiddleWare

import (
	"Golang/2025/01January/20250115/hertz-test/response"
	"Golang/2025/01January/20250115/hertz-test/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func Token() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		auth := ctx.Request.Header.Get("Authorization")

		if auth == "" {
			ctx.JSON(response.TokenErr, response.TokenError())
			ctx.Abort()
			return
		}

		claims, err := utils.VerifyJWT(auth)
		if err != nil {
			ctx.JSON(response.TokenErr, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set("GetName", claims)

		ctx.Next(c)
	}
}
