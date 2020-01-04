package Bai_1

import (
	"fmt"
	"time"
)

//khi noi den go routin va channal la dung toi phanmanh nhat cua go la concurrency

//cu phap khai bao
func Go_routin()  {
	//chay async
	go func() {
		fmt.Println("1")
	}()
	//chay asyn
	go func() {
		fmt.Printf("2")
	}()

	time.Sleep(time.Second)
}