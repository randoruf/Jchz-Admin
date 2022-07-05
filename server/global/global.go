package global

import (
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"jchz-admin/config"
)

var (
	JA_DB                  *gorm.DB                // 连接的数据库
	JA_CONFIG              config.ServerConfig     // 服务端设置
	JA_Concurrency_Control = &singleflight.Group{} // 并发控制
)
