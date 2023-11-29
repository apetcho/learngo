package functions

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func Functions() {
	fmt.Println()
	fmt.Println("-*-------------*-")
	fmt.Println("-*- Functions -*-")
	fmt.Println("-*-------------*-")
	result := plus(1, 2)
	fmt.Println("1 + 2 =", result)

	result = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", result)
}

// -*- MultipleReturnValue -*-
func values() (int, int) {
	return 3, 7
}

func MultipleReturnValue() {
	fmt.Println()
	fmt.Println("-*--------------------------*-")
	fmt.Println("-*- Multiple Return Values -*-")
	fmt.Println("-*--------------------------*-")
	a, b := values()
	fmt.Println(a)
	fmt.Println(b)

	_, c := values()
	fmt.Println(c)
}

// -*-
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(" ===> total =", total)
}

// -*- Variadics -*-
func Variadics() {
	fmt.Println()
	fmt.Println("-*---------------------*-")
	fmt.Println("-*- Variadic function -*-")
	fmt.Println("-*---------------------*-")
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

// -*- Closures -*-
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Closures() {
	fmt.Println()
	fmt.Println("-*------------*-")
	fmt.Println("-*- Closures -*-")
	fmt.Println("-*------------*-")
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
