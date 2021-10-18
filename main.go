package main

import (
	"fmt"
	"script-go-demo/jsonTest"
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
}
