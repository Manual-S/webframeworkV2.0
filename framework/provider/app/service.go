package app

import (
	"flag"

	"webframeworkV2.0/framework"
)

type HadeApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
}

func (h HadeApp) BaseFolser() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}

	var baseFolder string
	flag.StringVar(&baseFolder,
		"base_folder",
		"",
		"base_folder参数 默认当前路径")
	flag.Parse()

	if baseFolder != "" {
		return baseFolder
	}

	// todo
	return ""
}
