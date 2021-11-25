package main

import (
	"fmt"
	"runtime"
)

var sysos string

func init() {
	sysos = runtime.GOOS
}
func main() {
	fmt.Println("OS running on my system", sysos)
}
