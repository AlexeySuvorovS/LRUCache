package cache

import "errors"

type Key any
type Val any

type LRUI interface {
	Add(Key, Val)
	Get(Key) (Val, bool)
	Len() int
	GetQueue() []Key
}

func NewLRU(typeId string, cap int) (LRUI, error) {
	var ret LRUI
	var err error

	switch typeId {
	case "Slice", "slice":
		ret = NewLRUSlice(cap)
	case "List", "list":
		ret = NewLRUList(cap)
	default:
		err = errors.New("Undefined Type:" + typeId)

	}

	return ret, err
}
