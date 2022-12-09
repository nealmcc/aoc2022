package twod

import "fmt"

func ExamplePoint_Reduce() {
	a := Point{X: -12, Y: -3}
	b, scale := a.Reduce()
	fmt.Println(b, scale)
	// Output: {-4 -1} 3
}
