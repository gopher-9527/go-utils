package syncpool

import "sync"

func NewSyncPool() {
	syncPool := sync.Pool{}
	syncPool.New = func() interface{} {
		return nil
	}
}
