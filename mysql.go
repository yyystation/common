package common

import "go-micro.dev/v4/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

//获取mysql的配置

func GetMysqlFormConsul(config config.Config, path ...string) *MysqlConfig {
	MysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(MysqlConfig)
	return MysqlConfig
}
