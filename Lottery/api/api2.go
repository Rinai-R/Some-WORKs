package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.POST("/AddPrize", AddPrize)
	r.POST("/Host", Host)
	r.POST("/Lottery", QueryPrize)
	r.POST("/DelLottery", DelLottery)

	r.Run(":8080")
}
