package level2

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// -
func zeroptr(iptr *int) {
	*iptr = 0
}

// -*-
func Pointers() {
	fmt.Println()
	fmt.Println("-*------------*-")
	fmt.Println("-*- Pointers -*-")
	fmt.Println("-*------------*-")
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}
