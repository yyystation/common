package common

import (
	"strconv"

	"github.com/go-micro/plugins/v4/config/source/consul"
	"go-micro.dev/v4/config"
)

//设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		//设置配置中心的地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		//设置前缀，不设置前缀
		consul.WithPrefix(prefix),
		//是否移除前缀，这里是设置为true,表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}
	err = config.Load(consulSource)
	return config, err
}
