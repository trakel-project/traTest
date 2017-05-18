package data

type Result struct {
	Err    error
	Result string
}

const ContrastAddress string = "0x656cedbf993b54f5b0cf8262ea1d1a2c06f67711"
const URL string = "http://114.55.64.145:8081"

//存储数据的模型
type Contract struct {
	Funcname   string // 合约方法名
	Sender     string // 发送方地址
	Privatekey string // 发送方私钥
	Address    string // 合约地址
}

func (r *Result) SetResult(result string, err error) *Result {
	r.Err = err
	r.Result = result
	return r
}
