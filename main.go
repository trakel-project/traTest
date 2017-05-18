package main

import (
	"fmt"
	"traTest/data"
	"traTest/manager"
	"traTest/tester"
)

//合约对象
var mycontract = data.Contract{"getPassengerState", "0x3ddf839385969e9177804bf89778dd26401f033c", "0x6f50e13aa1bd3df3d7ea28896d10a69a5bbf6586d8a5e5c19ffefc9bb9b9146e", "0x656cedbf993b54f5b0cf8262ea1d1a2c06f67711"}

//一个测试者实例
var mytester = tester.Tester{10, 1000, 1, 2}

//管理者属于单例

func main() {
	fmt.Println(mycontract, mytester)
	manager.Test(mytester, mycontract)
}


