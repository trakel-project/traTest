package main

import (
	"log"
	"os"
	"time"

	"git.hyperchain.cn/yeyc/hyperkit/rpc"
)

const (
	passengerAddr   = "0x2cd84f9e3c182c5c543571ea00611c41009c7024"
	passengerPriKey = "0x437cace9ccb62f0e3e5bd71d2793aa8ac4a0e9d42262028e4a4dc7797d060dff"
	contractAddr    = "0x656cedbf993b54f5b0cf8262ea1d1a2c06f67711"
	contractABI     = `[{"constant":false,"inputs":[{"name":"x","type":"int256"},{"name":"y","type":"int256"},{"name":"time","type":"uint256"}],"name":"driverPickUpPassenger","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"int256"},{"name":"y","type":"int256"},{"name":"state","type":"bool"}],"name":"driverUpdatePos","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderDriver","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"time","type":"uint256"},{"name":"len","type":"uint256"},{"name":"x","type":"int256[]"},{"name":"y","type":"int256[]"}],"name":"driverFinishOrderWithRoute","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"s_x","type":"int256"},{"name":"s_y","type":"int256"},{"name":"d_x","type":"int256"},{"name":"d_y","type":"int256"}],"name":"calculatePreFeeWithPenatly","outputs":[{"name":"","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderPassInfo","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"getDriverRegiterState","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"driverCompetOrder","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderInfo0","outputs":[{"name":"id","type":"uint256"},{"name":"passenger","type":"address"},{"name":"s_x","type":"int256"},{"name":"s_y","type":"int256"},{"name":"d_x","type":"int256"},{"name":"d_y","type":"int256"},{"name":"distance","type":"int256"},{"name":"preFee","type":"int256"},{"name":"startTime","type":"uint256"},{"name":"passInfo","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderDrivInfo","outputs":[{"name":"drivInfo0","type":"string"},{"name":"drivInfo1","type":"string"},{"name":"drivInfo2","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"tsTotalNumOfDriver","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderActFee","outputs":[{"name":"","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"getAccountBalance","outputs":[{"name":"","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"isPassenger","type":"bool"}],"name":"getOrderID","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"getPassStateAndDriPos","outputs":[{"name":"state","type":"uint256"},{"name":"x","type":"int256"},{"name":"y","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"time","type":"uint256"}],"name":"driverFinishOrder","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"s_x","type":"int256"},{"name":"s_y","type":"int256"},{"name":"d_x","type":"int256"},{"name":"d_y","type":"int256"},{"name":"time","type":"uint256"},{"name":"passInfo","type":"string"},{"name":"sName","type":"string"},{"name":"dName","type":"string"}],"name":"passengerSubmitOrder","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"int256"},{"name":"y","type":"int256"},{"name":"threshold","type":"int256"}],"name":"getNearDrivers","outputs":[{"name":"","type":"uint256[5]"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"isPassenger","type":"bool"}],"name":"getDriverInfo","outputs":[{"name":"x","type":"int256"},{"name":"y","type":"int256"},{"name":"name","type":"address"},{"name":"info0","type":"string"},{"name":"info1","type":"string"},{"name":"info2","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderState","outputs":[{"name":"","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"int256"},{"name":"y","type":"int256"}],"name":"updatePassengerPos","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderDisAndActFee","outputs":[{"name":"distance","type":"int256"},{"name":"actFeeD","type":"int256"},{"name":"actFeeT","type":"int256"},{"name":"duration","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"cur_x","type":"int256"},{"name":"cur_y","type":"int256"}],"name":"driverCalculateActFee","outputs":[{"name":"","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderFeeStimeAndPlaceName","outputs":[{"name":"fee","type":"int256"},{"name":"time","type":"uint256"},{"name":"sName","type":"string"},{"name":"dName","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"newInfo0","type":"string"},{"name":"newInfo1","type":"string"},{"name":"newInfo2","type":"string"}],"name":"driverChangeInfo","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderInfo1","outputs":[{"name":"driver","type":"address"},{"name":"actFee","type":"int256"},{"name":"actFeeTime","type":"int256"},{"name":"pickTime","type":"uint256"},{"name":"endTime","type":"uint256"},{"name":"state","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"passengerPrepayFee","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderPreFee","outputs":[{"name":"","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"score","type":"int256"},{"name":"comment","type":"string"}],"name":"passengerJudge","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"getPassengerState","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"isPenalty","type":"bool"}],"name":"passengerCancelOrder","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"isPassenger","type":"bool"}],"name":"getJudge","outputs":[{"name":"avgScore","type":"int256"},{"name":"total","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"tsTotalNumOfOrder","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderStateAndDriverPos","outputs":[{"name":"state","type":"int256"},{"name":"x","type":"int256"},{"name":"y","type":"int256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"info0","type":"string"},{"name":"info1","type":"string"},{"name":"info2","type":"string"}],"name":"newDriverRegister","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"getDriverState","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"getDriverOrderPool","outputs":[{"name":"","type":"uint256[8]"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"orderIndex","type":"uint256"}],"name":"getOrderPlaceName","outputs":[{"name":"sName","type":"string"},{"name":"dName","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"driverIndex","type":"uint256"}],"name":"tsDriverInfo","outputs":[{"name":"isworking","type":"bool"},{"name":"x","type":"int256"},{"name":"y","type":"int256"},{"name":"state","type":"uint256"},{"name":"totalJudge","type":"int256"},{"name":"avgScore","type":"int256"},{"name":"name","type":"address"},{"name":"info0","type":"string"},{"name":"info1","type":"string"},{"name":"info2","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"driverCancelOrder","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"verify","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"inputs":[{"name":"cc","type":"address"}],"payable":false,"type":"constructor"}]`
)

func main() {
	log.Println("Connecting to Chain...")
	hrpc, err := rpc.NewRpc("http://114.55.64.145:8081", time.Second*10)
	if err != nil {
		log.Println("Connect to server fail", err)
		os.Exit(1)
	}

	//passengerSubmitOrder
	ret, err := hrpc.Invoke(passengerAddr, contractAddr, passengerPriKey, contractABI, "passengerSubmitOrder", true, "120143722", "30283618", "121000000", "31000000", "123", "pass0", "xixi", "yuquan")
	log.Printf("Invoke to \"passengerSubmitOrder\", Return Type: %T, Return Value: %v\n", ret, ret)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}