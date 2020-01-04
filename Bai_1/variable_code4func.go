package Bai_1

import "fmt"

func Code4func()  {
	var email = "lamn8@fpt.com"
	fullname:="nguyen lam"
	fmt.Println(email)
	fmt.Println(fullname)

	var number = 10
	var float = 10.234
	fmt.Println(number)
	fmt.Println(float)

	fmt.Printf("%T",float)
	fmt.Printf("%T",number)

}

func sum(a int, b int) int  {
	return  a+b

}

func sayhello(string  A) string  {
	return A
}
func Slice_arr()  {
	//slice la phan mo rong cua array nhung ko gioi hang cac phan tu
	s:= make([]string,0)
	// kieu slide s, co kieu du lieu la string, phan tu dau tien la rong
	//them phan tu vao slide dung append
	s =append(s, "a")
    s=append(s,"b")
    s=append(s,"c")
}
func Map_()  {
	_map := make(map[string]int)
	__map :=map[string]int{}


}

func _Array()  {
	//phai co gioi hang cac phan tu
	_arr := [4]int{}
	__arr:= [5] string{}
}

type (
	Person struct {
		Name string
		Age int
		Birthday string
	}
)

func Struct()  {
	p :=Person{
		Name:     "lamnn",
		Age:      10,
		Birthday: "02/09/1999",
	}
	fmt.Printf("%s",p.Name)
	name := p.Getname()
	fmt.Printf("%s",name)
}


//receiver function : khai bao kieu du lieu cho struct , giong method cua struct. gia tri ko bi thay doi,
func (p Person)Getname() string  {
	p.Name="ahihi"
	return p.Name
}
//receiver pointer : giong receiver function nhung du lei thay doi
func (p *Person)_GetName(_name string) string   {
	 p.Name = _name
	 return p.Name
}

