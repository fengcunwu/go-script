package panicTest

import "fmt"

// 注意：
// 1，recover 必须在 panic 发生之后收集
// 2，recover 在defer中执行时，需要加上一层匿名函数
func DealSome() {
	defer func () {
		if err := recover(); err != nil {
			fmt.Println("recover some error")
		}
		//fmt.Println("defer")
	}()
	panic("error")
}

func DealSomeV1() {
	if err := recover(); err != nil {
		fmt.Println("recover some error")
	}
	panic("error")
}

func DealSomeV2() {
	defer recover()
	panic("error")
}


func DealSomeV3() {
	panic("error")
	if err := recover(); err != nil{
		fmt.Println("recover some err")
	}

}
