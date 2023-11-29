package containers

import (
	"fmt"
	"slices"
)

func Arrays() {
	fmt.Println("-*-------------------------*-")
	fmt.Println("-*- Arrays data structure -*-")
	fmt.Println("-*-------------------------*-")
	var a [5]int
	fmt.Println("a:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

// -*- Slices -*-
func Slices() {
	fmt.Println("-*---------------*-")
	fmt.Println("-*- Slices data -*-")
	fmt.Println("-*---------------*-")
	var svec []string
	fmt.Println("uninit:", svec, svec == nil, len(svec) == 0)

	svec = make([]string, 3)
	fmt.Println("empty:", svec, "len:", len(svec), "cap:", cap(svec))

	svec[0] = "a"
	svec[1] = "b"
	svec[2] = "c"
	fmt.Println("set:", svec)
	fmt.Println("get:", svec[2])
	fmt.Println("len:", len(svec))

	svec = append(svec, "d")
	svec = append(svec, "e", "f")
	fmt.Println("append:", svec)

	cvec := make([]string, len(svec))
	copy(cvec, svec)
	fmt.Println("copy:", cvec)

	lslice := svec[2:5]
	fmt.Println("lslice:", lslice)

	lslice = svec[:5]
	fmt.Println("lslice2:", lslice)

	lslice = svec[2:]
	fmt.Println("lslice3:", lslice)

	tslice := []string{"g", "h", "i"}
	fmt.Println("dcl:", tslice)

	tslice2 := []string{"g", "h", "i"}
	if slices.Equal(tslice, tslice2) {
		fmt.Println("tslice == tslice2")
	}

	// -
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
