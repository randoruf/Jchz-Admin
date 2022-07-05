package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	configPath = "config.yaml"
)

type ServerConfig struct {
	SystemConfig SystemConfig `yaml:"System-Config"`
	TokenConfig  TokenConfig  `yaml:"Token-Config"`
	DbConfig     DbConfig     `yaml:"Db-Config"`
}

type SystemConfig struct {
	Port   string `json:"port" yaml:"port"`       // 端口
	DbType string `json:"db-type" yaml:"db-type"` // 数据库类型：默认 mysql （当前只支持 mysql）
}

type TokenConfig struct {
	AccessExpire uint64 `json:"access-expire" yaml:"access-expire"` // token 租期（单位：小时）
	AccessSecret string `json:"access-secret" yaml:"access-secret"` // jwt 密钥
	BufferTime   int64  `json:"buffer-time" yaml:"buffer-time"`     // jwt 续约缓冲时间
}

type DbConfig struct {
	Path         string `json:"path" yaml:"path"`                     // 服务器地址
	Port         string `json:"port" yaml:"port"`                     //:端口
	Config       string `json:"config" yaml:"config"`                 // 高级配置
	Dbname       string `json:"db-name" yaml:"db-name"`               // 数据库名
	Username     string `json:"username" yaml:"username"`             // 数据库用户名
	Password     string `json:"password" yaml:"password"`             // 数据库密码
	MaxIdleConns int    `json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}

type Mysql struct {
	DbConfig
}

// Dsn 获取 Database Source Name
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

// InitConfig 读取配置文件以初始化设置
func InitConfig() ServerConfig {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	var s ServerConfig
	err = yaml.Unmarshal(yamlFile, &s)
	return s
}
