package main

import "fmt"

type A interface {
	Run(string) string
}

type B interface {
	Say(string) int
}

type C struct {

}

func (c *C) Say(s string) int {
	fmt.Println(s)
	return 1
}

func (c *C) Run(a string) string {
	return a
}

var _ B = new(C)



func main()  {
	fmt.Println("hello")
}