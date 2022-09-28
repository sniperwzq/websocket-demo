package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"gitlab.zxmn2018.com/bigdata/go-zhise/single"
)

const (
	defaultName = "gz.config"
)

var gzConfPath string = "etc/websocket.yaml"

type Config struct {
	rest.RestConf
}

func LoadCfg(path string) *Config {
	gzConfPath = path
	return doGetOrLoadConf(gzConfPath)
}

func Cfg() *Config {
	return doGetOrLoadConf(gzConfPath)
}

func doGetOrLoadConf(path string) *Config {
	return single.GetOrSetFunc(defaultName, func() interface{} {
		var c Config
		conf.MustLoad(path, &c)
		return &c
	}).(*Config)
}
