package gorouting

import (
	"fmt"
	"math/rand"
	"time"
)

func F(a int)  {
	for i:=0;i<10;i++{
		fmt.Println(a ,":" ,i)
		amt:= time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond*amt)

	}
}