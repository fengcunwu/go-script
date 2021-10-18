package loadBalanceTest

import (
	"math/rand"
)

type LoadBalance struct {
	curIndex int
	addrList []string
}

func NewLoadBalance(addrList []string) *LoadBalance {
	return &LoadBalance{
		curIndex: 0,
		addrList: addrList,
	}
}

func (lb *LoadBalance)AppendAddr(addr string) {
	if lb != nil {
		lb.addrList = append(lb.addrList, addr)
	}
}

// GetAddrCycle 轮训负载均衡
func (lb *LoadBalance)GetAddrCycle() string {
	if lb != nil {
		temp := (lb.curIndex + len(lb.addrList)) % len(lb.addrList)
		lb.curIndex++
		return lb.addrList[temp]
	}
	return ""
}

// GetAddrCycle 轮训负载均衡
func (lb *LoadBalance)GetAddrRand() string {
	if lb != nil {
		// 返回一个随机数组， 先随机找一个数放在 数组末尾，在剩下的数中再找一个随机的数放在 n-1的末尾
		temp := rand.Perm(len(lb.addrList))
		return lb.addrList[temp[0]]
	}
	return ""
}


// 加权负载均衡

