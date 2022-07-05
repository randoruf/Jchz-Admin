package Initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jchz-admin/config"
	"jchz-admin/global"
	"strings"
)

// InitGorm 初始化 Gorm 连接数据库
func InitGorm() *gorm.DB {
	DbType := strings.ToLower(global.JA_CONFIG.SystemConfig.DbType)
	switch DbType {
	case "mysql":
		m := config.Mysql{global.JA_CONFIG.DbConfig}
		mysqlConfig := mysql.Config{
			DSN: m.Dsn(),
		}
		db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
		if err != nil {
			return nil
		} else {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxIdleConns(m.MaxIdleConns)
			sqlDB.SetMaxOpenConns(m.MaxOpenConns)
			return db
		}
	default:
		return nil
	}

}
