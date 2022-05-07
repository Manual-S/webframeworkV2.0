package gin

import "webframeworkV2.0/framework"

func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	// todo
	return nil
}

func (engine *Engine) IsBind(key string) bool {
	// todo
	return false
}

func (engine *Engine) Make(key string) (interface{}, error) {
	return nil, nil
}

func (engine *Engine) MustMake(key string) interface{} {
	return nil
}

func (engine *Engine) MakeNew(key string, params []interface{}) (interface{}, error) {
	// todo
	return nil, nil
}
