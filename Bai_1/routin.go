package Bai_1

import (
	"fmt"
	"time"
)

func f(from string)  {
	for i:=0;i< 3;i++ {
		fmt.Println(from,":",i)
	}
}
func Test_goroutin()  {
	f("direct")

	go f("goroutin")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}