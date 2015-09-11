/*
 Exercise: Slices
 Implement Pic.
 It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
 When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

 The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

 (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
 (Use uint8(intValue) to convert between types.)
*/
package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func main() {
	fmt.Println("### Start Slice Exercise ###")

	pic.Show(Pic)

	fmt.Println("### End Slice Exercise ###")
}

//dx and dy are length of slice
func Pic(dx, dy int) [][]uint8 {
	slice_dy := make([][]uint8, dy)

	for y := range slice_dy {
		slice_dy[y] = make([]uint8, dx)
		for x := range slice_dy[y] {
			v := (x*x + y*y) * 2
			//v := x*x + y*y
			//v := (x * x * y * y)
			slice_dy[y][x] = uint8(v)
		}
	}

	return slice_dy
}
