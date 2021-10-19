package main

import (
	"context"
	"fmt"
	"github.com/go-script/etcdTest"
	"strconv"
	"time"
)

var (
	endpoints = []string{"localhost:2379", "localhost:22379", "localhost:32379"}
	timeout = 1 * time.Second
)

func main () {

	client, err := etcdTest.NewEtcdClient(etcdTest.NewEtcdConfig(endpoints, timeout))
	if err != nil {
		panic(err)
	}
	defer client.EtcdClient.Close()
	for i := 0; i < 10; i++ {
		client.Put(context.Background(), "/sync-config/uve/generic_module", strconv.Itoa(i))
		fmt.Printf("put %d success\n", i)
		time.Sleep(2 * time.Second)
	}
}
