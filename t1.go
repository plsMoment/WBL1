package main

import "fmt"

type Human struct {
	name string
}

func (h *Human) GetName() string {
	return h.name
}

type Action struct {
	Human
}

func main() {
	alex := Human{"Alex"}
	ac := Action{alex}
	fmt.Printf("Name: %s\n", ac.GetName())
}
