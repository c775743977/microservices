package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql MysqlConfig
}

type MysqlConfig struct {
	DataSource string
}