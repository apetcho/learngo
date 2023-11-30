package level4

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"text/template"

	str "strings"
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

// -*----------------*-
// -*- (05) Recover -*-
// -*----------------*-
func mayPanic() {
	panic("a problem")
}

func Recover() {
	fmt.Println()
	fmt.Println("-*----------------*-")
	fmt.Println("-*- (05) Recover -*-")
	fmt.Println("-*----------------*-")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}

// -*------------------------*-
// -*- (06) StringFunctions -*-
// -*------------------------*-
var println = fmt.Println

func StringFunctions() {
	// fmt.Println()
	// fmt.Println("-*--------------*-")
	// fmt.Println("-*- (06) Defer -*-")
	// fmt.Println("-*--------------*-")
	header("(06) StringFunctions")
	println("Contains:  ", str.Contains("test", "es"))
	println("Count:     ", str.Count("test", "t"))
	println("HasPrefix: ", str.HasPrefix("test", "te"))
	println("HasSuffix: ", str.HasSuffix("test", "st"))
	println("Index:     ", str.Index("test", "e"))
	println("Join:      ", str.Join([]string{"a", "b", "c"}, "-"))
	println("Repeat:    ", str.Repeat("a", 5))
	println("Replace:   ", str.Replace("foo", "o", "0", -1))
	println("Repleace:  ", str.Replace("foo", "o", "0", 1))
	println("Split:     ", str.Split("a-b-c-d-e", "-"))
	println("ToLower:   ", str.ToLower("TEST"))
	println("ToUpper:   ", str.ToUpper("test"))
}

// -*- header(text) -*-
func header(text string) {
	num := len(text) + 2
	line := str.Join([]string{"-*--", str.Repeat("-", num), "--*-"}, "")
	fmt.Println()
	fmt.Println(line)
	fmt.Println("-*- ", text, " -*-")
	fmt.Println(line)
}

// -*-------------------------*-
// -*- (07) StringFormatting -*-
// -*-------------------------*-
type Point struct {
	x, y int
}

func StringFormatting() {
	header("(07) String formatting")
	point := Point{1, 2}
	fmt.Printf("struct1: %v\n", point)
	fmt.Printf("struct2: %+v\n", point)
	fmt.Printf("struct3: %#v\n", point)
	fmt.Printf("type: %T\n", point)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int: %d\n", 123)
	fmt.Printf("bin: %b\n", 14)
	fmt.Printf("char: %c\n", 33)
	fmt.Printf("hex: %x\n", 456)
	fmt.Printf("float1: %f\n", 78.9)
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)
	fmt.Printf("str1: %s\n", "\"string\"")
	fmt.Printf("str2: %q\n", "\"string\"")
	fmt.Printf("str3: %x\n", "hex this")
	fmt.Printf("pointer: %p\n", &point)
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	text := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(text)
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}

// -*----------------------*-
// -*- (08) TextTemplates -*-
// -*----------------------*-
func TextTemplates() {
	header("(08) Text Templates")

	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value: {{.}}\n"))
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C++", "C"})

	create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t2 := create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct{ Name string }{"Jane Doe"})
	t2.Execute(os.Stdout, map[string]string{"Name": "Mickey Mouse"})

	t3 := create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#", "C", "Python"})
}
