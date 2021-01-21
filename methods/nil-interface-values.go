// +build OMIT

package main

import (
	"fmt"
)

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// Error because there is no concrete method to call.
	//i.M()
}

func describe(i I) {
	fmt.Printf("(%v %T)\n", i, i)
}
