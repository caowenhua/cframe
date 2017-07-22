package sequence

import (
	"sync"
)

type MemorySequence struct {
	Data map[string]uint64
	lock sync.RWMutex
}

//query one sequence by id
func (mg *MemorySequence) Query(id string) (uint64, error) {
	mg.lock.RLock()
	var value = mg.Data[id] + 1
	mg.Data[id] = value
	mg.lock.RUnlock()
	return value, nil
}

func (mg *MemorySequence) Init() {
}

func (mg *MemorySequence) Save() {
}
