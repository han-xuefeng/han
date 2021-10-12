package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

type myInt int

func main()  {
	userType := reflect.Indirect(reflect.ValueOf(&User{Name: "jack", Age: 20}))

	a := "this is test"
	var b myInt = 34
	c := 25
	fmt.Println(reflect.TypeOf(a))

	fmt.Println(reflect.TypeOf(b))

	//fmt.Println(reflect.Indirect(reflect.TypeOf(&c)))

	cType := reflect.ValueOf(c)
	fmt.Println(cType.Float())

	fmt.Println(userType)

}