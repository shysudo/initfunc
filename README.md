# init func

# introduction

    Golang has special, niladic function init() for special purpouse of form a specific state for application.

    init() function runs a piece of code before running other part of the package, which allow application to be initialize in a specific state, such as have a specific configuration or set of resource with application need to start.

    init() function it is also used when importing a side effect, a technique used to set the state of the application by importing a package.

## declaring init()   

    Before declaring an init() function, lets see a function without init() function

## maig.go //refer the package withoutimport
   
        package main

        import "fmt"

        var sysos string

        func main() {
	        fmt.Println("OS running on my system", os)
        }

In this program we declared a global variable called sysos, default value of sysos is empty string

Lets run the program
    go run main.go

## output
    OS running on my system
    
Because sysos variable is black, we can see only the "OS running on my system" in output above.

Now we can fill the blank variable by introuducing an init() function that initialize the value of sysos to the os of system

## maig.go //refer the package withimport
   
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

In this program we assinged runtime.GOOS(imported runtime package) to sysos variable inside init() function. at runtime, sysos will be holding a value of the current system OS name

## output
        OS running on my system windows
        
Currently this program running on windows OS, so we can see in output "OS running on my system windows"

This illustrate that how init() works.

Another typical use case of init() when importing a package

## a.go // refere to packageimport package
        package packageimprt

        import "fmt"

        var AState string

        func init() {
        	AState = "WelCome in file a.go"
        	fmt.Println(AState)
        }

## b.go // refere to packageimport package
        package packageimprt

        import (
	        "fmt"
        )

        var BState string

        func init() {
        	BState = "Resume in file b.go"
        	fmt.Println(BState)
        }

Above two file a.go and b.go are in a package named "packageimport" and having a global variable AState, BState in a.go and b.go respectively, these variables are initialized to random string in init() function in both the file.

Now we have another package called "mainimport" that uses/calls "packageimport" package for accessing the globle variable AState and BState. 

## main.go //refer "mainimport" package

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
            fmt.Println("Main func")
            fmt.Println(Main)
        }

## output
        WelCome in file a.go
        Resume in file b.go
        Init func inside main package
        Main func
        Running from main.go

From the output of the we can figure out the order of init() func execution/call. In main.go file the package "github.com/shysudo/initfunc/packageimport" is imported, so init() function from this imported package execute/runs and runs init() function from main.go.

The order of the init() func execution/run based on the alphabetical order of the files in the imported package i.e init() function in a.go execute first "WelCome in file a.go" and init() function in b.go execute next "Resume in file b.go"