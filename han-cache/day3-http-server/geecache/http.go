package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultBasePath = "/_geecache"

// httpPool
type HTTPPool struct {
	self string
	basePath string
}

func NewHTTPPool (self string) *HTTPPool {
	return &HTTPPool{
		self: self,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool)Log(format string, v ...interface{})  {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}

func (p *HTTPPool)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if (strings.EqualFold(r.URL.Path, "/favicon.ico")) {
		w.Write([]byte(" "))
	}
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path: " + r.URL.Path)
	}

	p.Log("%s %s", r.Method, r.URL.Path)
	//fmt.Println(strings.Split(r.URL.Path[len(p.basePath):], "/"))
	parts := strings.SplitN(r.URL.Path[len(p.basePath)+1:], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]

	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}