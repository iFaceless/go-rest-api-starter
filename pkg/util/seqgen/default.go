// Package seqgen 简易发号器服务
// TODO: 将来迁移发号器服务，独立部署，对外提供 HTTP/RPC 接口
package seqgen

import (
	"os"
	"sync"

	"github.com/spf13/cast"
)

var (
	defaultGenerator *SeqGenerator
	lock             sync.Mutex
)

func NextID() (int64, error) {
	return getDefaultGenerator().NextID()
}

func getDefaultGenerator() *SeqGenerator {
	if defaultGenerator != nil {
		return defaultGenerator
	}

	roomID := os.Getenv("SEQ_SERVER_ROOM_ID")
	nodeID := os.Getenv("SEQ_SERVER_NODE_ID")

	if roomID == "" || nodeID == "" {
		panic("missing env SEQ_SERVER_ROOM_ID or SEQ_SERVER_NODE_ID")
	}

	lock.Lock()
	defer lock.Unlock()

	if defaultGenerator != nil {
		return defaultGenerator
	}

	sg, err := NewSeqGenerator(cast.ToInt(roomID), cast.ToInt(nodeID))
	if err != nil {
		panic(err)
	}

	defaultGenerator = sg
	return sg
}
