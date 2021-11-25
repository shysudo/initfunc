package main

import (
	"fmt"

	_ "github.com/shysudo/initfunc/packageimport"
)

var Main string

func init() {
	Main = "Running from main.go"
	fmt.Println("Init func inside main package")
}

func main() {
	fmt.Println(Main)
}
