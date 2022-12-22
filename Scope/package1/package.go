package package1

import "fmt"

var privateVar = "I am private" // not accesible outside
var PublicVar = "I am exported" // accesible outside, exported because of capitalization

func notExported() {

}

func Exported() {
	fmt.Println("I am exported function")
}
