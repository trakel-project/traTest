package manager

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"traTest/data"
	"traTest/tester"
)

var reconnectingtime int = 0

func Test(atester tester.Tester, contract data.Contract) {
	num := atester.Count
	log.Println("num is", num)
	interval := time.Duration(atester.Interval)
	log.Println("interval is", interval)
	group := atester.Group
	log.Println("group is", group)
	//先连接,连接失败重连
	hrpc, err := atester.Connect()
	for (err != nil) && (reconnectingtime < 14) {
		log.Panicln("Reconnecting...")
		reconnectingtime++
		hrpc, err = atester.Connect()
	}
	for currentGroup := 0; currentGroup < group; currentGroup++ {
		counterBad := 0
		counterGood := 0
		var mutex sync.Mutex
		var waitgroup sync.WaitGroup
		waitgroup.Add(num)
		for currentTest := 0; currentTest < num; currentTest++ {
			go func() {
				defer waitgroup.Done()
				start := time.Now()
				result := atester.InvokeContract(hrpc, contract)
				duration := time.Since(start)
				log.Println(start, duration)
				if result.Err != nil {
					mutex.Lock()
					counterBad++
					mutex.Unlock()
					return
				}
				if duration > time.Duration(atester.Timeout)*time.Second {
					mutex.Lock()
					counterBad++
					mutex.Unlock()
				} else {
					mutex.Lock()
					counterGood++
					mutex.Unlock()
				}
			}()
		}
		waitgroup.Wait()
		connectToFile(contract.Funcname, currentGroup, counterGood, counterBad)
		time.Sleep(time.Second * interval)
	}
}
func writeToFile(f *os.File, context string) {
	f.Write([]byte(context))
}

//还应该记录调用的方法，以及调用的时间
func connectToFile(funcName string, currentGroup int, counterGood int, counterBad int) (err error) {
	f, err := os.OpenFile("./result.txt", os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		log.Println("File open failed", err)
		return
	}
	if currentGroup == 0 {
		writeToFile(f, fmt.Sprintf("Invoke %s\n", funcName))
	}
	log.Println(fmt.Sprintf("No.%d Good has %d Bad has %d, Good rate is %.2f%%\n", currentGroup+1, counterGood, counterBad, float64(counterGood)/float64(counterBad+counterGood)*100))
	writeToFile(f, fmt.Sprintf("No.%d Good has %d Bad has %d, Good rate is %.2f%%\n", currentGroup+1, counterGood, counterBad, float64(counterGood)/float64(counterBad+counterGood)*100))
	if currentGroup == 9 {
		writeToFile(f, "\n")
	}
	return
}

//把数据写入文件
