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

<<<<<<< HEAD
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
=======
	f, err := os.OpenFile("./result.txt", os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		log.Println("File open failed", err)
		os.Exit(1)
	}

	var num int
	fmt.Println("input times")
	fmt.Scanf("%d", &num)

	str := fmt.Sprintf("Input size: %d\n", num)
	f.Write([]byte(str))
	str = fmt.Sprintln("Start Time", time.Now())
	fmt.Print(str)
	f.Write([]byte(str))
	for j := 0; j < 10; j++ {
		counterBad := 0
		counterGood := 0
		timeout := time.Second * 10
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(num)
		//fmt.Println("Start Time", time.Now())
		for i := 0; i < num; i++ {
			//用于模拟在1秒内等间隔发送num笔订单
			//如果模拟的是同时发送num笔订单则注释下面语句
			time.Sleep(time.Second / time.Duration(num))
			//ii := i
			go func() {
				defer wg.Done()
				start := time.Now()
				_, err := hrpc.Invoke(passengerAddr, contractAddr, passengerPriKey, contractABI, "passengerSubmitOrder", false, "120143722", "30283618", "121000000", "31000000", "123", "passinfo", "qidian", "zhongdian")
				duration := time.Since(start)
				if err != nil {
					//fmt.Println("no.", ii, err)
					mu.Lock()
					counterBad++
					fmt.Printf("\rNo.%d Good has %d Bad has %d, Good rate is %.2f%%", j+1, counterGood, counterBad, float64(counterGood)/float64(counterBad+counterGood)*100)
					mu.Unlock()
					return
				}
				if duration > timeout {
					mu.Lock()
					counterBad++
					fmt.Printf("\rNo.%d Good has %d Bad has %d, Good rate is %.2f%%", j+1, counterGood, counterBad, float64(counterGood)/float64(counterBad+counterGood)*100)
					mu.Unlock()
				} else {
					mu.Lock()
					counterGood++
					fmt.Printf("\rNo.%d Good has %d Bad has %d, Good rate is %.2f%%", j+1, counterGood, counterBad, float64(counterGood)/float64(counterBad+counterGood)*100)
					mu.Unlock()
					//fmt.Println("no.", ii, duration.Seconds())
				}
			}()
		}
		// fmt.Println("End Time", time.Now())
		wg.Wait()
		str = fmt.Sprintf("No.%d Good has %d Bad has %d, Good rate is %.2f%%\n", j+1, counterGood, counterBad, float64(counterGood)/float64(counterBad+counterGood)*100)
		f.Write([]byte(str))
		fmt.Println()
		time.Sleep(time.Second * 3)
>>>>>>> b76dd4e26be52a263904bea263b126cbb3504272
	}

	//TODO 测试query

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
