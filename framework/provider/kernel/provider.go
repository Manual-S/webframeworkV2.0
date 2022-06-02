package kernel

import (
	"webframeworkV2.0/framework"
	"webframeworkV2.0/framework/gin"
)

// HadeKernelProvider 服务的提供者
type HadeKernelProvider struct {
	HttpEngine *gin.Engine
}

func (provider *HadeKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeKernelService
}
