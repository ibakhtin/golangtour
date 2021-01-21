// +build OMIT

package main

import "fmt"

type I interface {
	M()
}

type T1 struct {
	S string
}

func (t T1) M() {
	fmt.Println(t.S)
}

type T2 int

func (t T2) M() {
	fmt.Println(t)
}

func main() {
	var i I

	i = T1{"hello"}
	i.M()

	i = T2(42)
	i.M()
}
