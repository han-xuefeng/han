package singleflight

import "sync"

type call struct {
	wg sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait() //如果有请求就等待
		return c.val, c.err //请就结束，返回结果
	}

	c := new(call)
	c.wg.Add(1)  //发起请求  加锁
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn() //调用查询
	c.wg.Done() //关闭锁

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
