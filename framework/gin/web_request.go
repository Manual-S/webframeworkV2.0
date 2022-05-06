package gin

import "github.com/spf13/cast"

type IRequest interface {

	// 请求url中带参数的地址 比如
	// http://localhost:8080/hello?id=14444

	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	//DefaultQueryFloat64(key string, def float64) (float64, bool)
	//DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	//DefaultQueryStringSlice(key string, def []string) ([]string, bool)
}

func (ctx *Context) QueryAll() map[string][]string {
	// todo 这里gin框架有一个cache
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

func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}

	}

	return def, false
}

func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}

	}

	return def, false
}

func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}

	return def, false
}
