package packageimprt

import (
	"fmt"
)

var BState string

func init() {
	BState = "Resume in file b.go"
	fmt.Println(BState)
}
