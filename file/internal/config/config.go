package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	// 新增
	Path string `json:",default=."`
}
