package app

import (
	"errors"

	"webframeworkV2.0/framework"
)

// HadeAppProvider 提供APP的具体实现方法
type HadeAppProvider struct {
	BaseFolder string
}

// NewHadeApp 初始化HadeApp
func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("params error")
	}

	container := params[0].(framework.Container)
	baseFolder := params[1].(string)

	return &HadeApp{
		baseFolder: baseFolder,
		container:  container,
	}, nil
}
