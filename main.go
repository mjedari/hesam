package main

import (
	"fmt"
	"github.com/mjedari/hesam/pkg"
)

func main() {
	fmt.Println("Hello World!")
	time := pkg.GetRealTime()
	fmt.Println("Time is: ", time)
}
