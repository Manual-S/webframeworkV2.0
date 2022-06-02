package kernel

import (
	"net/http"

	"webframeworkV2.0/framework/gin"
)

type HadeKernelService struct {
	engine *gin.Engine
}

// NewHadeKernelService 初始化web引擎服务实例
func NewHadeKernelService(params ...interface{}) (interface{}, error) {
	engine := params[0].(*gin.Engine)
	return HadeKernelService{
		engine: engine,
	}, nil
}

// HttpEngine 返回web引擎服务实例
func (s *HadeKernelService) HttpEngine() http.Handler {
	return s.engine
}
