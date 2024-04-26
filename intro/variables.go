package main

import "fmt"

// inferred constant in the global namespace
// const = immutable
const (
    iteration = 1
    keyLen = 10
)
// in the global namespace, go can perform type inference
// normal variable initialization
// inferred variable value (go default string value: "")
var (
    name             = "Foo"
    firstName string = "Foo"
    lastName  string
)

var x = 10

func main() {
    // local variables of the main function scope
    version := 1            // infer integer
    otherVersion := "Bar"   // infer string
    anotherVersion := 10.1  //infer float32

    var version int         // default 0-initialized
     
    fmt.Println(version)
}
