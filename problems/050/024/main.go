package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	listAll(2)
	listAll(3)

	listNth(3, 2)

	start := time.Now()
	listNth(10, 1000000)
	log.Println("cost", time.Since(start))
}

func listNth(placeNum, nth int) {
	log.Println("list", placeNum, "permutations", nth, "th")
	if placeNum <= 0 || placeNum > 10 {
		log.Fatal("placeNum must be in [1, 10]")
	}

	var per Permutation = make([]string, placeNum)
	for i := 0; i < placeNum; i++ {
		per[i] = fmt.Sprintf("%d", i)
	}

	total := multi(changeLimits(per))

	if nth > total {
		log.Fatalf("nth %d out of range", nth)
	}

	changed, err := nthPermutation(per, nth)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\t%dth %v", nth, changed)
}

func listAll(placeNum int) {
	log.Println("list all", placeNum, "permutations")
	if placeNum <= 0 || placeNum > 10 {
		log.Fatal("placeNum must be in [1, 10]")
	}

	var per Permutation = make([]string, placeNum)
	for i := 0; i < placeNum; i++ {
		per[i] = fmt.Sprintf("%d", i)
	}

	total := multi(changeLimits(per))

	for i := 1; i <= total; i++ {
		changed, err := nthPermutation(per, i)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\t%3dth %v", i, changed)
	}
}

func nthPermutation(origin Permutation, nth int) (Permutation, error) {
	if nth == 1 {
		return origin, nil
	}
	size := len(origin)
	limits := changeLimits(origin)
	// log.Printf("n=%d limit=%v", nth, limits)

	changeTime := nth - 1
	var changes []int = make([]int, size)
	for i := 0; i < size-1 && changeTime != 0; i++ {
		changeTimes := multi(limits[i+1:])
		if changeTimes > changeTime {
			// log.Printf("n=%d i=%d left=%d changeTimes=%d pass", nth, i, changeTime, changeTimes)
			continue
		}
		changes[i] = changeTime / changeTimes

		changeTime = changeTime - (changeTimes * changes[i])
		// log.Printf("n=%d i=%d left=%d changeTimes=%d change=%d %v", nth, i, changeTime, changeTimes, changes[i], limits[i+1:])
	}
	// log.Printf("n=%d changes=%v", nth, changes)

	cur := make(Permutation, size)
	copy(cur, origin)

	var err error
	for index, changes := range changes {
		cur, err = change(cur, index, changes)
		// log.Printf("n=%d change index=%d changes=%d cur=%v", nth, index, changes, cur)
		if err != nil {
			return nil, err
		}
	}

	return cur, nil
}

func multi(is []int) int {
	var sum int = 1
	for _, i := range is {
		sum *= i
	}
	return sum
}

func changeLimits(origin Permutation) []int {
	size := len(origin)
	limits := make([]int, size)
	for i := 0; i < size; i++ {
		limits[i] = size - i
	}
	return limits
}

// change i是数字位数-1，j是第几次变化
func change(p Permutation, index, changes int) (Permutation, error) {
	if index >= len(p) {
		return nil, fmt.Errorf("index %d out of range", index)
	}
	if changes >= len(p)-index {
		return nil, fmt.Errorf("index %d changes %d out of range", index, changes)
	}

	base := p[:index]
	left := p[index : index+changes]
	s := p[index+changes]
	var right Permutation
	if index+changes+1 < len(p) {
		right = p[index+changes+1:]
	}

	// log.Printf("base=%v left=%v s=%v right=%v", base, left, s, right)

	var output Permutation
	output = append(output, base...)
	output = append(output, s)
	// log.Printf("output=%v s%v", output, s)
	output = append(output, left...)
	// log.Printf("output=%v %v", output, left)
	output = append(output, right...)
	// log.Printf("output=%v %v", output, right)

	return output, nil
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

type Permutation []string

func (p Permutation) String() string {
	return strings.Join(p, "")
}
