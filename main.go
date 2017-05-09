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

	num := 300
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		time.Sleep(time.Millisecond * 3)
		ii := i
		go func() {
			start := time.Now()
			defer wg.Done()
			_, err := hrpc.Invoke(passengerAddr, contractAddr, passengerPriKey, contractABI, "passengerSubmitOrder", false, "120143722", "30283618", "121000000", "31000000", "123", "passinfo", "qidian", "zhongdian")
			//log.Printf("Invoke to \"passengerSubmitOrder\", Return Type: %T, Return Value: %v\n", ret, ret)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			fmt.Println("no.", ii, time.Since(start).Seconds())
		}()
	}
	wg.Wait()

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
