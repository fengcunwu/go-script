package reflectTest

import (
	"fmt"
	"reflect"
)

// 本文件主要用于 测试 标准库reflect包中的相关函数，


type Person struct {
	name string
}

func (p *Person) Run () {
	if p != nil {
		fmt.Printf("%s is running\n", p.name)
	}
}

func NewPerson (name string) *Person {
	return &Person{name: name}
}

func (p *Person)ExecuteFunction(funcName string) {
	vaule := reflect.ValueOf(p)
	f := vaule.MethodByName("Run")
	f.Call([]reflect.Value{})
}
