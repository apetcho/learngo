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
