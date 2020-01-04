package main

import (
	"fmt"
	"log"
	"time"
)

//func input()(int ,error){
//	var n int
//	fmt.Print("nhap mot so ko am")
//	_,err :=fmt.Scanf("%d",&n)
//	if err == nil && n<0 {
//		return n,fmt.Errorf("loi %d l so nguyn am",n)
//	}
//	return n,err
//}
//func main()  {
//	   input()
//
//}



func main()  {
	a:=1*time.Minute
	fmt.Println(a)
	deadline:=time.Now().Add(a)
	s:=time.Now().Before(deadline)
	fmt.Println(time.Now())
	fmt.Println(deadline)
	fmt.Println(s)
	log.Printf("server ko tra loi")
}