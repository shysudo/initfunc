package packageimprt

import "fmt"

var AState string

func init() {
	AState = "WelCome in file a.go"
	fmt.Println(AState)
}
