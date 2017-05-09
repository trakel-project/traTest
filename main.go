package main

import (
	"log"
	"os"
	"time"
	"traTest/constABI"

	"fmt"
	"sync"

	"git.hyperchain.cn/yeyc/hyperkit/rpc"
)

var contractABI = constABI.ABI

const (
	passengerAddr   = "0x3ddf839385969e9177804bf89778dd26401f033c"
	passengerPriKey = "0x6f50e13aa1bd3df3d7ea28896d10a69a5bbf6586d8a5e5c19ffefc9bb9b9146e"
	contractAddr    = "0x656cedbf993b54f5b0cf8262ea1d1a2c06f67711"
)

func main() {
	log.Println("Connecting to Chain...")
	hrpc, err := rpc.NewRpc("http://114.55.64.145:8081", time.Second*10)
	if err != nil {
		log.Println("Connect to server fail", err)
		os.Exit(1)
	}

	var num int
	fmt.Println("input times")
	fmt.Scanf("%d", &num)
	counterBad := 0
	counterGood := 0
	timeout := time.Second * 10
	var mug sync.Mutex
	var mub sync.Mutex
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		//time.Sleep(time.Millisecond * 3)
		ii := i
		go func() {
			defer wg.Done()
			start := time.Now()
			_, err := hrpc.Invoke(passengerAddr, contractAddr, passengerPriKey, contractABI, "passengerSubmitOrder", false, "120143722", "30283618", "121000000", "31000000", "123", "passinfo", "qidian", "zhongdian")
			duration := time.Since(start)
			if err != nil {
				fmt.Println("no.", ii, err)
				mub.Lock()
				counterBad++
				mub.Unlock()
				return
			}
			if duration > timeout {
				mub.Lock()
				counterBad++
				mub.Unlock()
			} else {
				mug.Lock()
				counterGood++
				mug.Unlock()
				fmt.Println("no.", ii, duration.Seconds())
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Good has %d\n Bad has %d\n", counterGood, counterBad)
	//passengerSubmitOrder
	// ret, err := hrpc.Invoke(passengerAddr, contractAddr, passengerPriKey, contractABI, "passengerSubmitOrder", false, "120143722", "30283618", "121000000", "31000000", "123", "passinfo", "qidian", "zhongdian")
	// log.Printf("Invoke to \"passengerSubmitOrder\", Return Type: %T, Return Value: %v\n", ret, ret)
	// if err != nil {
	// 	log.Println(err)
	// 	os.Exit(1)
	// }
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for {
	// 		time.Sleep(time.Second * 1.0)
	// 		ret, err := hrpc.Invoke(passengerAddr, contractAddr, passengerPriKey, contractABI, "getPassengerState", true)
	// 		log.Printf("Invoke to \"getPassengerState\", Return Type: %T, Return Value: %v\n", ret, ret)
	// 		if err != nil {
	// 			log.Println(err)
	// 			os.Exit(1)
	// 		}
	// 		currentPassengerState := ret.(*big.Int).Uint64()
	// 		if currentPassengerState == 2 {
	// 			break
	// 		}
	// 	}
	// }()
	// wg.Wait()
}
