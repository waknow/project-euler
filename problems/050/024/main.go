package main

import (
	"fmt"
	"strings"
)

func main() {
	const s = "0123456789"
	size := len(s)
	is := make([]int, size)
	for i := range is {
		is[i] = factorial(size - i)
	}

	want := 1000000

	for i, v := range is {
		if v > want {
			fmt.Println(size-i, 0)
			continue
		}
		n := want / v
		fmt.Println(size-i, n)
		want = want % v
	}
	fmt.Println("want:", want)

	fmt.Println("1st", nthPermutation(s, 2, 2))

}

// nthPermutation i是数字位数-1，j是第几次变化
func nthPermutation(s string, i, j int) string {
	index := len(s) - i - 1
	if index >= len(s) {
		panic("i >= len(s)")
	}

	var ss []string
	for _, v := range s {
		ss = append(ss, string(v))
	}

	s1 := ss[index]
	ss[index] = ss[index+j]
	ss[index+j] = s1
	return strings.Join(ss, "")
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
