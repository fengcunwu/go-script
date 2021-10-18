package interview

import (
	"fmt"
	"sync"
)


// 1， 统计字符串数组中每一个字符出现的次数（同步和并发版本）
////==============////==============////==============////==============////==============////==============////==============
type Set map[rune]int
// CountRune 统计字符串数组中每个字符出现的个数
func CountRune(strs []string) Set {
	result := make(map[rune]int)
	for _, str := range strs {
		for _, char := range str {
			result[char]++
		}
	}
	return result
}

// CountRuneAsync 异步统计字符串数组中每个字符串的格式，count表示可以开启的线程数目
func CountRuneAsync(strs []string, count int) Set {
	ch := make(chan map[rune]int, count)
	result := make(chan map[rune]int, 1)
	temp := make(map[rune]int)
	var wg sync.WaitGroup

	for _, str := range strs {
		wg.Add(1)
		go Producer(&wg, ch, str)
	}

	go Consumer(ch, result)
	wg.Wait()
	close(ch)

	select {
	case res := <- result:
		temp = res
	}
	return temp
}

func Producer(wg *sync.WaitGroup, ch chan map[rune]int, str string) {
	result := make(map[rune]int)
	for _, char := range str{
		result[char]++
	}
	ch <- result
	wg.Done()
}

func Consumer(ch chan map[rune]int, result chan map[rune]int) {
	res := make(map[rune]int)
	for {
		value, ok := <-ch
		if ok {
			for k, v := range value {
				res[k] += v
			}
		} else {
			break
		}
	}
	result <- res
}

//func main() {
//	ch := make(chan map[rune]int, 2)
//
//	result := make(chan map[rune]int, 1)
//
//
//	wg := sync.WaitGroup{}
//
//	for _, str := range []string{"aaa", "abc", "ddd", "ddddd"} {
//		wg.Add(1)
//		go Producer(&wg, ch, str)
//	}
//	go Consumer(ch, result)
//
//	wg.Wait()
//	close(ch)
//
//	select {
//	case res := <- result:
//		log.Println(res)
//	}
//}


// 2, 三个线程按顺序打印 cat dog pig 10次
////==============////==============////==============////==============////==============////==============////==============

// Print 多个线程 错误示例
func Print() {
	catChan := make(chan string, 1)
	dogChan := make(chan string, 1)
	pigChan := make(chan string, 1)
	var wg sync.WaitGroup

	catChan <- "cat"
	for i := 0; i < 10; i ++ {
		wg.Add(3)
		go print(&wg, catChan, dogChan, "dog")
		go print(&wg, dogChan, pigChan, "pig")
		go print(&wg, pigChan, catChan, "cat")
	}
	wg.Wait()
}

func print(wg *sync.WaitGroup, ch1 chan string, ch2 chan string, input string) {
	fmt.Println(<-ch1)
	ch2 <- input
	wg.Done()
}

func printThree (str string, preCh, nextCh chan int) {
	for {
		temp := <-preCh
		if temp == -1 {
			nextCh <- temp
			return
		}
		fmt.Println(str)
		nextCh <- temp
	}
}

func PrintThree() {
	catCh := make(chan int)
	dogCh := make(chan int)
	pigCh := make(chan int)
	endCh := make(chan int)

	go printThree("cat", catCh, dogCh)
	go printThree("dog", dogCh, pigCh)
	go printThree("pig", pigCh, endCh)

	for i := 0; i < 3; i++ {
		catCh <- 1
		<- endCh
	}
}
