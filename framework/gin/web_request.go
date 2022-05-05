package gin

import "github.com/spf13/cast"

type IRequest interface {
	DefaultQueryInt(key string, def int) (int, bool)
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.Request != nil {
		return map[string][]string(ctx.Request.URL.Query())
	}

	return map[string][]string{}
}

func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}

	}

	return def, false
}
