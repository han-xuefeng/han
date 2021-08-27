package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req *http.Request
	// re
	Path string
	Method string
	Params map[string]string
	StatusCode int
	handlers []HandlerFunc
	index int
	engine *Engine
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
		index: -1,
	}
}

func (c *Context) Next () {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// 获取post参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 获取get
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 设置状态吗
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//设置响应头部
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 响应字符串
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// json
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data
func (c *Context)Data(code int, data []byte)  {
	c.Status(code)
	c.Writer.Write(data)
}

//html
func (c *Context)HTML(code int, name string, data interface{})  {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	//c.Writer.Write([]byte(html))
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}

func (c *Context)Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}