package geecache

import (
	"fmt"
	"geecache/consistenthash"
	pb "geecache/geecachepb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	defaultBasePath = "/_geecache"
	defaultReplicas = 50
)

// http 客户端
type httpGetter struct {
	baseURL string
}
// httpPool http服务端
type HTTPPool struct {
	self string
	basePath string
	mu sync.Mutex
	peers *consistenthash.Map //服务端的缓存节点管理
	httpGetters map[string]*httpGetter  // 服务端的缓存真实节点
}

func (p *HTTPPool) Get(in *pb.Request, out *pb.Response) error {
	return nil
}

func NewHTTPPool (self string) *HTTPPool {
	return &HTTPPool{
		self: self,
		basePath: defaultBasePath,
	}
}

// 服务端设置节点
func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock() // 设置节点加锁，防止数据不一致
	defer p.mu.Unlock()
	p.peers = consistenthash.New(defaultReplicas, nil)
	p.peers.Add(peers...)

	// 把http的服务端的缓存真实节点存到httpGetters
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}
// 服务端选择节点功能
func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

var _ PeerPicker = (*HTTPPool)(nil)

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

	// Write the value to the response body as a proto message.
	body, err := proto.Marshal(&pb.Response{Value: view.ByteSlice()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(body)
}

//从group中查找缓存
func (h *httpGetter) Get(in *pb.Request, out *pb.Response) error {
	u := fmt.Sprintf(
		"%v%v/%v",
		h.baseURL,
		url.QueryEscape(in.Group),
		url.QueryEscape(in.Key),
	)
	res, err := http.Get(u)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned: %v", res.Status)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}
	if err = proto.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("decoding response body: %v", err)
	}
	return nil
}
var _ PeerGetter = (*httpGetter)(nil)