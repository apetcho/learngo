package level2

import (
	"errors"
	"fmt"
	"math"
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

// -*--------------*-
// -*- Interfaces -*-
// -*--------------*-
type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(geom geometry) {
	fmt.Println(geom)
	fmt.Println(geom.area())
	fmt.Println(geom.perim())
}

func Interfaces() {
	fmt.Println()
	fmt.Println("-*--------------*-")
	fmt.Println("-*- Interfaces -*-")
	fmt.Println("-*--------------*-")

	rect := rectangle{width: 3, height: 4}
	circ := circle{radius: 5}
	measure(rect)
	measure(circ)
}

// -*-------------------*-
// -*- StructEmbedding -*-
// -*-------------------*-
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func StructEmbedding() {
	fmt.Println()
	fmt.Println("-*--------------------*-")
	fmt.Println("-*- Struct embedding -*-")
	fmt.Println("-*--------------------*-")
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}

// -*------------*-
// -*- Generics -*-
// -*------------*-
func MapKeys[K comparable, V any](dict map[K]V) []K {
	result := make([]K, 0, len(dict))
	for k := range dict {
		result = append(result, k)
	}
	return result
}

// -
type List[T any] struct {
	head, tail *element[T]
}

// -
type element[T any] struct {
	next *element[T]
	val  T
}

// -
func (lst *List[T]) Push(val T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: val}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: val}
		lst.tail = lst.tail.next
	}
}

// -
func (lst *List[T]) GetAll() []T {
	var result []T
	for elem := lst.head; elem != nil; elem = elem.next {
		result = append(result, elem.val)
	}
	return result
}

// -*-
func Generics() {
	fmt.Println()
	fmt.Println("-*------------*-")
	fmt.Println("-*- Generics -*-")
	fmt.Println("-*------------*-")
	var dict = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(dict))

	_ = MapKeys[int, string](dict)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}

// -*-----------------*-
// -*- ErrorHandling -*-
// -*-----------------*-
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (err *argError) Error() string {
	return fmt.Sprintf("%d - %s", err.arg, err.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// -
func ErrorHandling() {
	fmt.Println()
	fmt.Println("-*------------------*-")
	fmt.Println("-*- Error Handling -*-")
	fmt.Println("-*------------------*-")
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
