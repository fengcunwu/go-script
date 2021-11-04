package Cache

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache  map[int]*list.Element
	ll  *list.List
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache: make(map[int]*list.Element),
		ll: list.New(),
	}
}


func (this *LRUCache) Get(key int) int {
	if this == nil {
		return -1
	}
	v, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.ll.MoveToFront(v)
	result, _ := v.Value.(*entry)
	return result.key
}


func (this *LRUCache) Put(key int, value int)  {
	if this == nil {
		return
	}

	v, ok := this.cache[key]

	if ok {
		v.Value = &entry{
			key: key,
			value: value,
		}
		this.ll.MoveToFront(v)
	} else {
		if len(this.cache) >= this.capacity {
			r := this.ll.Back()
			if r != nil {
				res := r.Value.(*entry)
				delete(this.cache, res.key)
				this.ll.Remove(r)
			}
		}
		e := &list.Element{
			Value: &entry{
				key: key,
				value: value,
			},
		}

		this.cache[key] = e
		this.ll.PushFront(e)
		this.capacity += 1
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	ca := Constructor(2)
	ca.Put(1, 1)
	ca.Put(2, 2)
	ca.Get(1)

	ca.Print()
}
func (c LRUCache)Print() {
	head := c.ll.Front()

	for head.Next() != nil {
		fmt.Println("--->", head.Value.(*entry))
		head = head.Next()
	}
}
