package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()

	r.POST("/register", register)                         // 注册
	r.POST("/login", login)                               // 登录
	r.POST("/AlterPassword", AlterPassword)               // 改密码
	r.POST("/DeleteUser", DeleteUser)                     //删除用户
	r.POST("/OnlineChange", JWTMiddleware(), OnlineAlter) //不要密码，直接更改，需要验证jwt

	r.Run(":8088") // 跑在 8088 端口上
}
