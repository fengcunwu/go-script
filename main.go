package main

import (
	"container/list"
	"fmt"
	"github.com/go-script/jsonTest"
	"math"
	"sync"
	"time"
)

//var (
//	endpoints = []string{"localhost:2379", "localhost:22379", "localhost:32379"}
//	timeout = 1 * time.Second
//)

func main () {

	//client, err := etcdTest.NewEtcdClient(etcdTest.NewEtcdConfig(endpoints, timeout))
	//if err != nil {
	//	panic(err)
	//}
	//defer client.EtcdClient.Close()
	//fmt.Println("start watch")
	//for {
	//	client.Watch(context.Background(), "test-key")
	//}

	//temp := interview.CountRune([]string{"addd", "ddd", "a"})
	//temp2 := interview.CountRuneAsync([]string{"addd", "ddd", "a", "fffff"}, 4)
	//log.Println(temp)
	//log.Println(temp2)


	//jsonTest.Parse()


	ch := make(chan int)
	var wg sync.WaitGroup
	//go func() {
	//	for {
	//		temp := <- ch
	//		if temp == 0 {
	//			break
	//		}
	//		log.Println(temp)
	//	}
	//	wg.Done()
	//} ()

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("stop")
				return
			default:
				time.Sleep(time.Millisecond * 50)
				fmt.Println("go on")
			}
		}
	} ()

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("stop")
				return
			default:
				time.Sleep(time.Millisecond * 50)
				fmt.Println("go on")
			}
		}
	} ()

	//ch <- 1
	//ch <- 1

	close(ch)
	wg.Wait()
	time.Sleep(time.Second * 2)
	//interview.PrintThree()
	//jsonTest.Parse()
	var per = &jsonTest.Person{Name: "dddd"}
	var cur = &jsonTest.Person{
		Name: "wfc",
	}
	var result = []*jsonTest.Person{per, cur}

	for index, res := range result {
		if index == 1 {
			res.Name = "1"
		} else {
			res.Name = "2"
		}
	}


	for _, res := range result {
		fmt.Println(res.Name)
	}

	fmt.Println("======result")

	//ca := Constructor(1)
	//ca.Put(2, 1)
	//ca.Get(2)
	//fmt.Println(len(ca.cache))
	//fmt.Println(ca.capacity)
	//
	//ca.Print()
	//fmt.Println(ca.Get(2))

	var nums = []int{3,2,3,1,2,4,5,5,6}
	fmt.Println("====", findKthLargest(nums, 4))
	fmt.Println("=====", nums)
}


func coinChange(coins []int, amount int) int {
	dp := make([]int64, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt64
	}
	for _, coin := range coins {
		for i := 0; i <= amount; i++ {
			if coin <= i {
				dp[i] = min(dp[i], dp[i - coin] + 1)
			}
		}
	}
	fmt.Println(dp[amount])
	fmt.Println(dp[amount] == math.MaxInt64)

	if dp[amount] != math.MaxInt64 {
		return int(dp[amount])
	} else {
		return -1
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}


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
	return result.value
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
		e := &entry{
			key: key,
			value: value,
		}

		this.cache[key] = this.ll.PushFront(e)
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func (c LRUCache)Print() {
	head := c.ll.Front()

	for head != nil {
		fmt.Print(head.Value.(*entry), "--->")
		head = head.Next()
	}
}


func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}
	quickSort(nums, 0, len(nums))
	return nums[len(nums) - k]
}

func quickSort(nums []int, left,right int) {
	if left < right {
		ret := deal(nums, left, right)
		fmt.Println("------", nums)
		quickSort(nums, left, ret)
		quickSort(nums, ret + 1, right)
	}
}

func deal(nums []int, left, right int) int {
	begin, end, key := left, right-1, nums[right - 1]
	for begin < end {
		for begin < end && nums[begin] <=  key {
			begin++
		}

		for begin < end && nums[end] >= key {
			end--
		}
		if begin != end {
			swap(nums, begin, end)
		}
	}

	if begin != right {
		swap(nums, begin, right-1)
	}
	return begin
}


func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

