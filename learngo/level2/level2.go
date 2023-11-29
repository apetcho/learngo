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
