package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func main()  {
	value := reflect.ValueOf(&User{})
	fmt.Printf("%T,%V", value,value)
	modelType := reflect.Indirect(value).Type()

	fmt.Println(modelType)
}