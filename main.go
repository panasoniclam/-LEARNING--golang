package main

import "fmt"

type Dog struct {
    Name string
    Id int
    saidNumber int
}
type Animal interface {
    Said() string
    SaidNumber() int
}
type Getting func(name string) string

func (d *Dog)Said()  string{
    d.saidNumber ++
    return fmt.Sprintf("%s gau gau",d.Name)
}

func (g Getting)Exclumation(name string) string  {
    return  name + "!"
}

var english = Getting(func(name string) string{
    return  name + "hello"
})

func main()  {
    fmt.Println(english.Exclumation("lamnn"))
}