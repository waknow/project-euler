package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	const wantSize = 1000

	f1 := make(BigUint, wantSize)
	f2 := make(BigUint, wantSize)
	f1[len(f1)-1], f2[len(f2)-1] = 1, 1

	for i := 2; ; i++ {
		temp := f2.Copy()
		f2 = f2.Add(f1)
		f1 = temp
		// fmt.Printf("%4d: %3d - %s\n", i+1, f2.Len(), f2)

		if f2.Len() >= wantSize {
			fmt.Println("Found!", i+1, f2.Len())
			break
		}
	}
	fmt.Println("Elapsed:", time.Since(start))
}

type BigUint []int8

func (left BigUint) Add(right BigUint) BigUint {
	maxLen := len(left)
	if len(right) > maxLen {
		maxLen = len(right)
	}

	result := make(BigUint, maxLen+1)

	leftIndex := len(left) - 1
	rightIndex := len(right) - 1
	resultIndex := len(result) - 1
	var carry int8

	for ; resultIndex >= 0; resultIndex-- {
		if leftIndex < 0 && rightIndex < 0 {
			break
		}

		var leftDigit, rightDigit int8
		if leftIndex >= 0 {
			leftDigit = left[leftIndex]
		}
		if rightIndex >= 0 {
			rightDigit = right[rightIndex]
		}

		temp := leftDigit + rightDigit + carry
		// fmt.Printf("%-2d+%-2d=%2d: %d+%d+%d=%-2d result=%-2d carry=%-2d\n", leftIndex, rightIndex, resultIndex, leftDigit, rightDigit, carry, temp, temp%10, temp/10)
		result[resultIndex] = temp % 10
		carry = temp / 10

		rightIndex--
		leftIndex--
	}
	if carry != 0 {
		result[resultIndex] = carry
		// fmt.Printf("result[%d]=%d\n", resultIndex, result[resultIndex])
	}
	return result
}

func (left BigUint) Copy() BigUint {
	result := make(BigUint, len(left))
	copy(result, left)
	return result
}

func (left BigUint) Len() int {
	isPrefixZero := true
	var count int
	for _, v := range left {
		if isPrefixZero && v == 0 {
			continue
		}
		isPrefixZero = false
		count++
	}
	return count
}

func (left BigUint) String() string {
	buf := bytes.NewBuffer(nil)

	isPrefix := true
	for _, a := range left {
		if a == 0 && isPrefix {
			continue
		}
		isPrefix = false
		buf.WriteString(string(rune(48 + a)))
	}
	if buf.Len() == 0 {
		buf.WriteString("0")
	}

	return buf.String()
}
