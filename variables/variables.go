package main

import (
	"fmt"
	"reflect"
)

var (
	// name, course string
	// module       float64
	// name, course, module = "Nigel", "Docker Deep Dive", 3.2
	name   = "Nigel"
	course = "Docker Deep Dive"
	module = 3.2
)

func main() {
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
}
