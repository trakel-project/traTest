package tester

import (
	"log"
	//	"os"
	"time"
	"traTest/constABI"
	"traTest/data"

	"git.hyperchain.cn/yeyc/hyperkit/rpc"
)

var ABI = constABI.ABI

//专门测试合约
type Tester struct {
	Group    int // 测多少组
	Count    int // 并发量
	Interval int // 组间延时 毫秒
	Timeout  int // 超时 秒
}


//连接方法，返回结果

func (tester *Tester) Connect() (*rpc.Rpc, error) {
	//建立连接
	hrpc, err := rpc.NewRpc(data.URL, time.Duration(tester.Timeout)*time.Second)
	if err != nil {
		log.Println("Connect to server fail", err)
	}
	log.Println("Connecting to Chain...")
	//写文件交给manager做
	return hrpc, err
}

//调用合约，参数是合约对象，返回调用结果

func (tester *Tester) InvokeContract(hrpc *rpc.Rpc, contract data.Contract) data.Result {
	//如果用指针作为返回值，会有空指针的问题
	var result data.Result

	_, err := hrpc.Invoke(contract.Sender, contract.Address, contract.Privatekey, ABI, contract.Funcname, false) //正式调合约方法
	//	return result.SetResult("Invoking is over",err) 这里不对
	if err != nil {
		result = *result.SetResult("Invoke contract failed", err)
		log.Println("Invoke contract failed", err)
		result.Err = err
		return result
	}
//	log.Println("Success")
	result.Result = "success"
	return result
}
