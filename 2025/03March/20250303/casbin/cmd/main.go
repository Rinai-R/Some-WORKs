package main

import (
	"Golang/2025/03March/20250303/casbin/Initialize"
	"Golang/2025/03March/20250303/casbin/global"
	"Golang/2025/03March/20250303/casbin/router"
)

func main() {
	Initialize.InitDB()
	Initialize.InitCasbin(global.DB)
	router.InitRouter()
}
