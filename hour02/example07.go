package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string = "string"
	var i int = 10
	var f float32 = 1.2

	fmt.Println(reflect.TypeOf(s), s)
	fmt.Println(reflect.TypeOf(i), i)
	fmt.Println(reflect.TypeOf(f), f)
}
