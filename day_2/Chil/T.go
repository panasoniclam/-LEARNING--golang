package child

import (
	"errors"
	"fmt"
)

type (
	Dog struct {
		Name string
		saidNumber int
		Id int
	}
	Cat struct {
		Name string
		saidNumber int
		Id int

	}
	Animal interface {
		Said() string
		SaidNumber() int
	}
)

func (c *Cat) Said() string{
	c.saidNumber ++
	return fmt.Sprintf("%s mew mew", c.Name)
}

func (d *Dog)Said() string{
	d.saidNumber++
	return  fmt.Sprintf("%s gau gau", d.Name)
}
func (c *Cat)SaidNumber() int  {
	return c.saidNumber
}
func (d *Dog)SaidNumber() int {
	return  d.saidNumber
}
func NewAnimal(animalType string, name string) (A Animal, error error)  {
	if animalType == "dog" {
		return  &Dog{
			Name:       name,
			Id:         0,
		}, nil
	}
	if animalType == "cat" {
		return &Cat{
			Name:       name,

			Id:         0,
		}, nil
	}
	return nil, errors.New(" unsupport animal type")
}