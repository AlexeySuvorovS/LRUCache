package cache

import (
	"strconv"
	"testing"
)

type createStruct struct {
	TypeId     string
	ExpectedOK bool
}

var lruCreate = []createStruct{
	{TypeId: "slice", ExpectedOK: true},
	{TypeId: "list", ExpectedOK: true},
	{TypeId: "anyOther", ExpectedOK: false},
}

type hitStruct struct {
	key Key
	val Val
	exp bool
}

var hitCheck = []hitStruct{
	{key: 1, val: "abc1", exp: false},
	{key: 99, val: "abc99", exp: true},
	{key: 100, val: "abc100", exp: true},
	{key: 101, val: "abc101", exp: true},
}

func TestCreate(t *testing.T) {
	for _, v := range lruCreate {

		_, err := NewLRU(v.TypeId, 2)

		if err != nil {
			if v.ExpectedOK == true {
				t.Fatalf("Expeceted OK (%s), but error (%s)", v.TypeId, err.Error())
			}
		} else {
			if v.ExpectedOK == false {
				t.Fatalf("Expeceted Not OK (%s), but error id NIL", v.TypeId)
			}
		}

	}
}

func TestCapacitySlice(t *testing.T) {
	lenChecked := 100

	lru, err := NewLRU("slice", lenChecked)

	if err != nil {
		t.Fatal("Create Error")
	}

	for i := 1; i < 1000; i++ {
		val := "abc" + strconv.Itoa(i)
		lru.Add(i, val)
	}

	if lru.Len() != lenChecked {
		t.Fatalf("Slice Len expected %d, but get %d", lenChecked, lru.Len())
	}
}

func TestCapacityList(t *testing.T) {
	lenChecked := 100

	lru, err := NewLRU("list", lenChecked)

	if err != nil {
		t.Fatal("Create Error")
	}

	for i := 1; i < lenChecked*2; i++ {
		val := "abc" + strconv.Itoa(i)
		lru.Add(i, val)
	}

	if lru.Len() != lenChecked {
		t.Fatalf("List Len expected %d, but get %d", lenChecked, lru.Len())
	}
}

func TestGetSlice(t *testing.T) {
	lenChecked := 100

	lru, err := NewLRU("slice", lenChecked)

	if err != nil {
		t.Fatal("Create Error")
	}

	for i := 1; i <= lenChecked+1; i++ {
		val := "abc" + strconv.Itoa(i)
		lru.Add(i, val)
	}

	for _, v := range hitCheck {
		val, ok := lru.Get(v.key)

		if ok != v.exp && val != v.val {
			t.Fatalf("HIT: Expected %s, but get %s for %d -> %s", v.val, val, v.key, v.val)
		}

		if ok != v.exp {
			t.Fatalf("HIT: Expected %v, but get %v for %d -> %s", v.exp, ok, v.key, v.val)
		}
	}

}

func TestGetList(t *testing.T) {
	lenChecked := 100

	lru, err := NewLRU("list", lenChecked)

	if err != nil {
		t.Fatal("Create Error")
	}

	for i := 1; i <= lenChecked+1; i++ {
		val := "abc" + strconv.Itoa(i)
		lru.Add(i, val)
	}

	for _, v := range hitCheck {
		val, ok := lru.Get(v.key)

		if ok != v.exp && val != v.val {
			t.Fatalf("HIT Val: Expected %s, but get %s for %d -> %s", v.val, val, v.key, v.val)
		}

		if ok != v.exp {
			t.Fatalf("HIT: Expected %v, but get %v for %d -> %s", v.exp, ok, v.key, v.val)
		}
	}

}

func BenchmarkSlice(b *testing.B) {
	for j := 0; j < b.N; j++ {

		lenChecked := 1000

		lru, err := NewLRU("list", lenChecked)

		if err != nil {
			b.Fatal("Create Error")
		}

		for i := 1; i <= lenChecked+1; i++ {
			val := "abc" + strconv.Itoa(i)
			lru.Add(i, val)
		}

		for _, v := range hitCheck {
			_, ok := lru.Get(v.key)

			if ok != v.exp {
				b.Fatalf("HIT: Expected %v, but get %v for %d -> %s", v.exp, ok, v.key, v.val)
			}
		}
	}
}

func BenchmarkList(b *testing.B) {
	for j := 0; j < b.N; j++ {

		lenChecked := 1000

		lru, err := NewLRU("list", lenChecked)

		if err != nil {
			b.Fatal("Create Error")
		}

		for i := 1; i <= lenChecked+1; i++ {
			val := "abc" + strconv.Itoa(i)
			lru.Add(i, val)
		}

		for _, v := range hitCheck {
			_, ok := lru.Get(v.key)

			if ok != v.exp {
				b.Fatalf("HIT: Expected %v, but get %v for %d -> %s", v.exp, ok, v.key, v.val)
			}
		}
	}
}
