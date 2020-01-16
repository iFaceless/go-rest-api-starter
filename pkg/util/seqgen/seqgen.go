// Package seqgen 是一个基于 Twitter Snowflake 算法实现的发号器。
//
// |-|------------------|-------------------|---------------------|
// |0|timestamp(41 bits)|worker id (10 bits)|sequence id (12 bits)|
// |-|------------------|-------------------|---------------------|
//
// 这里需要注意的是，时间戳是最大可以表示 2^41-1 毫秒，也就是到 2039-09-07 23:47:35
// 所以，这里可以考虑减去最近的一个时间戳，来换取更长远的发号时间。
// worker id 表示的是机器编号，编号规则可以考虑如下：
// <机房 ID 2 bits><机器 ID 8 bits>
// 如此，可以保证最多支持部署在四个机房，每个机房最多部署 255 个节点。
//
// 参考文章：https://segmentfault.com/a/1190000014767902
//
package seqgen

import (
	"fmt"
	"sync"
	"time"
)

// epoch 时间戳统一减去该值（2019-10-13 22:54:00.000），一旦服务运行后，记得不要修改这个值了
// 否则可能会导致生成重复的 id
const epoch = 1570978440000

const (
	bitNumberOfTimestamp    = 41
	bitNumberOfServerRoomID = 2
	bitNumberOfServerNodeID = 8
	bitNumberOfSequence     = 12

	timestampBitShift = bitNumberOfSequence + bitNumberOfServerNodeID + bitNumberOfServerRoomID
)

var (
	maxValueOfTimestamp    = (1 << bitNumberOfTimestamp) - 1
	maxValueOfServerRoomID = (1 << bitNumberOfServerRoomID) - 1
	maxValueOfServerNodeID = (1 << bitNumberOfServerNodeID) - 1
	maxValueOfSequence     = (1 << bitNumberOfSequence) - 1
)

var (
	errTimestampExceedsLimit    = fmt.Errorf("invalid timestamp, exceeds the limit %d", maxValueOfTimestamp)
	errServerRoomIDExceedsLimit = fmt.Errorf("invalid server room id, exceeds the limit %d", maxValueOfServerRoomID)
	errServerNodeIDExceedsLimit = fmt.Errorf("invalid server node id, exceeds the limit %d", maxValueOfServerNodeID)
)

type SeqGenerator struct {
	workerID int
	seq      int
	ts       int64
	lock     sync.Mutex
}

func NewSeqGenerator(serverRoomID, serverNodeID int) (*SeqGenerator, error) {
	if serverRoomID > maxValueOfServerRoomID {
		return nil, errServerRoomIDExceedsLimit
	}

	if serverNodeID > maxValueOfServerNodeID {
		return nil, errServerNodeIDExceedsLimit
	}

	workerID := (serverRoomID << bitNumberOfServerNodeID) | serverNodeID
	return &SeqGenerator{workerID: workerID}, nil
}

func (sg *SeqGenerator) String() string {
	return fmt.Sprintf("SeqGenerator(wid=%d)", sg.workerID)
}

func (sg *SeqGenerator) NextID() (int64, error) {
	sg.lock.Lock()

	validCurrentTimestamp := func() int64 {
		ts := time.Now().UTC().UnixNano() / 1e6
		if ts > int64(maxValueOfTimestamp) {
			sg.lock.Unlock()
			panic(errTimestampExceedsLimit)
		}
		return ts
	}

	now := validCurrentTimestamp()
	if now == sg.ts {
		// 同一毫秒内，seq++
		sg.seq++
		if sg.seq > maxValueOfSequence {
			// 等待 1 毫秒
			for now <= sg.ts {
				now = validCurrentTimestamp()
			}

			sg.seq = 0
			sg.ts = now
		}
	} else {
		sg.seq = 0
		sg.ts = now
	}

	sg.lock.Unlock()
	id := (int64(now-epoch) << (timestampBitShift)) | int64(sg.workerID<<bitNumberOfSequence) | int64(sg.seq)
	return id, nil
}
