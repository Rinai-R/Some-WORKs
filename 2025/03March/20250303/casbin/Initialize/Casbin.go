package Initialize

import (
	"Golang/2025/03March/20250303/casbin/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"log"
)

func InitCasbin(db *gorm.DB) {
	// 创建适配器
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatal("failed to create adapter:", err)
	}

	// 加载模型
	global.Casbin, err = casbin.NewEnforcer("./2025/03March/20250303/casbin/global/conf/model.conf", adapter)
	if err != nil {
		log.Fatal("failed to create enforcer:", err)
	}

	// 加载策略
	err = global.Casbin.LoadPolicy()
	if err != nil {
		log.Fatal("failed to load policy:", err)
	}
}
