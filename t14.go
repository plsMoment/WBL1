package main

import (
	"fmt"
	"reflect"
)

// typeof returns string that determine type of parameter v.
// Detects types: int, string, bool and any chan
func typeof(v any) string {
	switch reflect.TypeOf(v).Kind().String() {
	case "int":
		return "int"
	case "string":
		return "string"
	case "bool":
		return "bool"
	case "chan":
		return "channel"
	default:
		return ""
	}
}

func main() {
	a := 1
	str := "aboba"
	f := true
	var ch chan bool
	fmt.Println(typeof(a))
	fmt.Println(typeof(str))
	fmt.Println(typeof(f))
	fmt.Println(typeof(ch))
}
