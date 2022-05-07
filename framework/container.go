// Package framework 服务容器
package framework

import (
	"errors"
	"sync"
)

type Container interface {
	// Bind 绑定一个服务提供者 如果关键字凭证已经存在 会进行替换操作 返回error
	Bind(provider ServiceProvider) error

	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务
	Make(key string) (interface{}, error)

	MustMake(key string) interface{}

	// MakeNew 根据关键字凭证获取一个服务
	MakeNew(key string, params []interface{}) (interface{}, error)
}

type WebContainer struct {
	Container

	// providers 存储注册的服务提供者 key为字符串凭证
	providers map[string]ServiceProvider

	// instances 存储具体的实例 key为字符串凭证
	instances map[string]interface{}

	// lock用于锁住对容器的变更操作
	lock sync.RWMutex
}

func NewWebContainer() Container {
	return &WebContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// Bind 将服务容器与关键字做绑定
func (w *WebContainer) Bind(provider ServiceProvider) error {
	w.lock.Lock()
	defer w.lock.Unlock()

	key := provider.Name()

	w.providers[key] = provider

	if provider.IsDefer() == false {
		// 判断是否在绑定的时候进行实例化

	}

	return nil
}

// Make 提供获取服务实例
func (w *WebContainer) Make(key string) (interface{}, error) {
	return w.make(key, nil, false)
}

func (w *WebContainer) findServiceProvider(key string) ServiceProvider {
	w.lock.Lock()
	defer w.lock.Unlock()

	if sp, ok := w.providers[key]; ok {
		// 已经注册了服务的提供者
		return sp
	}

	return nil
}

func (w *WebContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(w); err != nil {
		return nil, err
	}

	method := sp.Register(w)
	ins, err := method(w)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (w *WebContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	sp := w.findServiceProvider(key)
	if sp == nil {
		// 没有提供服务的注册者 报错
		return nil, errors.New(key + "have not register")
	}

	if ins, ok := w.instances[key]; ok {
		// 说明容器中已经实例化过
		// 直接使用实例化过的对象即可
		return ins, nil
	}

	// 说明容器中没有实例化过 则进行一次实例化
	inst, err := w.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	w.instances[key] = inst

	return inst, nil
}

func (w *WebContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	// todo
	return nil, nil
}
