package geecache

import pb "geecache/geecachepb"

type PeerPicker interface {
	// 根据key 获取对应的节点
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	// 从对应group中查询缓存值
	Get(in *pb.Request, out *pb.Response) (error)
}