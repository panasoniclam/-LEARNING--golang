package main

import (
	"fmt"
	child "github.com/panasoniclam/golang-basic/day_2/Chil"
	"os"
)

func main()  {
	 a, err := child.NewAnimal("dog","lucky")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(a.Said())
	 fmt.Println(a.Said())
      fmt.Printf("animal said %d time(s)", a.SaidNumber())
}