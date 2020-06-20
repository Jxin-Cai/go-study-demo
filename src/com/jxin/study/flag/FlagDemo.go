package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("jxin", "", "test")

func main() {
	flag.Parse()
	fmt.Println("value: %s", *mode)
	fmt.Println("point: %p", mode)
}
