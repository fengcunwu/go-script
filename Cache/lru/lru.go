package lru

import (
	"container/list"
	"github.com/go-script/Cache"
)

type lru struct {
	maxBytes int // 最大容量
	useBytes int // 已使用字节数
	list *list.List // go 内置双向链表
	cache map[string]*list.Element // 缓存数据，key对应的双链表的节点
}

type item struct {
	key string
	value interface{}
}

func NewCache(maxBytes int) Cache.Cache {
	return &lru{
		maxBytes: maxBytes,
		useBytes: 0,
		list:     list.New(),
		cache:     make(map[string]*list.Element),
	}
}

func (l *lru) Set(key string, value interface{}) {
	if element, ok := l.cache[key]; ok { // 之前存过， 不涉及容量判断删除元素
		l.list.MoveToFront(element) // 移动到 对头
		element.Value = value
		return
	}

	for l.maxBytes != 0 && len(l.cache) >= l.maxBytes {
		l.list.Remove(l.list.Back()) // 删除最后一个节点
	}

	element := &item{
		key:   key,
		value: value,
	}

	e := l.list.PushFront(element)
	l.cache[key] = e
}

func (l *lru) Get(key string) interface{} {
	return nil
}


func (l *lru) Del (key string) {

}

func (l *lru) Len () int{
	return 0
}

func (l *lru) UseBytes() int {
	return 0
}
