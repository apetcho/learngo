package level2

import (
	"fmt"
	"unicode/utf8"
)

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

// -*- StringAndRune -*-
func StringAndRune() {
	fmt.Println()
	fmt.Println("-*-------------------*-")
	fmt.Println("-*- String and Rune -*-")
	fmt.Println("-*-------------------*-")
	const s = "สวัสดี"
	fmt.Println("len(s):", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		// -
		examineRune(runeValue)
	}
}

// -
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

// -*-----------*-
// -*- Structs -*-
// -*-----------*-
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func Structs() {
	fmt.Println()
	fmt.Println("-*----------*-")
	fmt.Println("-*- Struct -*-")
	fmt.Println("-*----------*-")
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("John"))

	sean := person{name: "Sean", age: 50}
	fmt.Println(sean.name)
	seanp := &sean
	fmt.Println(seanp.age)
	seanp.age = 51
	fmt.Println(seanp.age)

	// -
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

// -*-----------*-
// -*- Methods -*-
// -*-----------*-
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func Methods() {
	fmt.Println()
	fmt.Println("-*-----------*-")
	fmt.Println("-*- Methods -*-")
	fmt.Println("-*-----------*-")
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
