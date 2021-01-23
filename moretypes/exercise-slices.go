// +build OMIT

package main

import (
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	// Create picture
	pic := make([][]uint8, dx)
	for i := range pic {
		pic[i] = make([]uint8, dy)
	}

	// Fill picture
	for i := range pic {
		for j := range pic[i] {
			pic[i][j] = uint8(i * j)
		}
	}
	return pic
}

func Show(p [][]uint8) {
	for i := range p {
		for j := range p[i] {
			fmt.Printf("%3x", p[i][j])
		}
		fmt.Println()
	}
}

func main() {
	Show(Pic(12, 12))
}
