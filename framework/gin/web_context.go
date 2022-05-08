package gin

import (
	"context"

	"webframeworkV2.0/framework"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

// Bind 实现container的绑定封装
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

// context 实现container的几个封装

func (ctx *Context) Make(key string) (interface{}, error) {
	return ctx.container.Make(key)
}

func (ctx *Context) MustMake(key string) interface{} {
	return ctx.container.MustMake(key)
}

func (ctx *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	// todo
	return nil, nil
}
