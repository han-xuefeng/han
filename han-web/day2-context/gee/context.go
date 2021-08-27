package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req *http.Request
	Path string
	Method string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}
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
func (c *Context)HTML(code int, html string)  {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}