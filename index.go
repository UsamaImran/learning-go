package main

import (
	"fmt"
	"math"
	"strconv"
)

func main (){
	fmt.Print("before")
	fmt.Println("Hello World")
	fmt.Println("testing logger")
	
	var value int 
	fmt.Print(value)

	common := "common variable"
	common = "asdasdasd"

	numberImplicit := 2300000032423423234
	fmt.Println(common)
	fmt.Println(numberImplicit)
	fmt.Print(common,numberImplicit,"\n","asdasdasd", value, math.Pow(2,3), strconv.Itoa(123))
}