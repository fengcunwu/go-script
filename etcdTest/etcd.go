package etcdTest

import (
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

type EtcdConfig struct {
	Endpoints []string
	DialTimeout time.Duration
}

type EtcdOp struct {
	EtcdClient    *clientv3.Client
}

func NewEtcdConfig(endPoints []string, timeout time.Duration) *EtcdConfig {
	return &EtcdConfig{
		Endpoints:   endPoints,
		DialTimeout: timeout,
	}
}


func NewEtcdClient(config *EtcdConfig) (*EtcdOp, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            config.Endpoints,
		DialTimeout:          config.DialTimeout,
	})
	if err != nil {
		return nil ,err
	}
	return &EtcdOp{EtcdClient: cli}, nil
}

func (e *EtcdOp) Put(ctx context.Context, key, value string){
	_, err := e.EtcdClient.Put(ctx, key, value)
	if err != nil {
		panic(err)
	}
}

func callBack(ctx context.Context, ev *clientv3.Event) {
	key := string(ev.Kv.Key)
	value := string(ev.Kv.Value)
	fmt.Println(key+"======"+value)
}

func (e *EtcdOp) Watch(ctx context.Context, key string) error {
	var watchCh clientv3.WatchChan = e.EtcdClient.Watch(ctx, key)


	for {
		select {
		case <-ctx.Done():
			return errors.New("ctx.Done")

		case resp, ok := <-watchCh:
			if !ok {
				watchCh = e.EtcdClient.Watch(ctx, key)
			} else {
				if resp.Canceled {
					watchCh = e.EtcdClient.Watch(ctx,key)
				} else {
					for _, ev := range resp.Events {
						callBack(ctx, ev)
					}
				}
			}
		}
	}
}
