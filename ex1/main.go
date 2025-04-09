package main

import "fmt"

// separating testable parts of the code, ie. the domain and the actual print, usually those are combined
func hello() string {
	return "hello world"
}
func main() {
	fmt.Println(hello())

}
