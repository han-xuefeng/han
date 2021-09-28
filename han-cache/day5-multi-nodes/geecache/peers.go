package geecache

type PeerPicker interface {
	// 根据key 获取对应的节点
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	// 从对应group中查询缓存值
	Get(group string, key string) ([]byte, error)
}