# init func

# introduction

    Golang has special, niladic function init() for special purpouse of form a specific state for application.

    init() function runs a piece of code before running other part of the package, which allow application to be initialize in a specific state, such as have a specific configuration or set of resource with application need to start.

    init() function it is also used when importing a side effect, a technique used to set the state of the application by importing a package.

# declaring init()   

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

Another typical use case of init() is to use when importing a package