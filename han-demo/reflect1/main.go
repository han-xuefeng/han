package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
}

func main()  {

	userValue := reflect.ValueOf(&User{Name: "jack", Age: 18})

	userValue = reflect.Indirect(userValue)
	fmt.Println(userValue.Method(0))
	//userType := userValue.Type()
	//fmt.Println(userType.Field(0))


}

