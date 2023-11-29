package basics

import (
	"fmt"
	"math"
)

func Hello() {
	fmt.Println("Hello World!")
}

func Values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func Variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

// -
const str string = "constant"

func Constant() {
	fmt.Println(str)
	const num = 50000000
	const dec = 3e20 / num
	fmt.Println(dec)
	fmt.Println(int64(dec))
	fmt.Println(math.Sin(num))
}

func For(){
	i := 1
	for i <= 3{
		fmt.Println(i)
		i = i + 1
	}

	// -
	for j := 7; j <= 9; j++{
		fmt.Println(j)
	}
	// -
	for{
		fmt.Println("loop")
		break
	}
	// -
	for n := 0; n <= 5; n++{
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
