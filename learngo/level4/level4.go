package level4

import (
	"cmp"
	"fmt"
	"os"
	"slices"
)

// -*----------------*-
// -*- (01) Sorting -*-
// -*----------------*-
func Sorting() {
	fmt.Println()
	fmt.Println("-*----------------*-")
	fmt.Println("-*- (01) Sorting -*-")
	fmt.Println("-*----------------*-")
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}

// -*---------------------------*-
// -*- (02) SortingByFunctions -*-
// -*---------------------------*-
func SortingByFunctions() {
	fmt.Println()
	fmt.Println("-*-----------------------------*-")
	fmt.Println("-*- (02) Sorting by functions -*-")
	fmt.Println("-*-----------------------------*-")

	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}

// -*---------------*-
// -*- (03) Panics -*-
// -*---------------*-
func Panics() {
	fmt.Println()
	fmt.Println("-*---------------*-")
	fmt.Println("-*- (03) Panics -*-")
	fmt.Println("-*---------------*-")
	panic("a problem")
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

// -*--------------*-
// -*- (04) GoDefer -*-
// -*--------------*-
func createFile(p string) *os.File {
	fmt.Println("creating")
	fp, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return fp
}

func writeFile(fp *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(fp, "data")
}

func closeFile(fp *os.File) {
	fmt.Println("closing")
	err := fp.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func GoDefer() {
	fmt.Println()
	fmt.Println("-*--------------*-")
	fmt.Println("-*- (04) Defer -*-")
	fmt.Println("-*--------------*-")

	fp := createFile("/tmp/defer.txt")
	defer closeFile(fp)
	writeFile(fp)
}
