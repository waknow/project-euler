package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
	"time"
)

var enableLog = true

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	start := time.Now()
	defer func() {
		fmt.Println("elapse", time.Since(start))
	}()

	const limit uint = 28123

	abundantsLog := fileLogln(ctx, "number-abundants.txt")
	prefectsLog := fileLogln(ctx, "number-perfects.txt")
	deficientsLog := fileLogln(ctx, "number-deficients.txt")

	var abundants []uint
	for i := uint(1); i <= limit; i++ {
		is := divisors(i)
		total := sum(is)
		if total > i {
			abundants = append(abundants, i)
			abundantsLog(i, "<", total, "=", SumExpr(is))
		} else if total == i {
			prefectsLog(i, "=", total, "=", SumExpr(is))
		} else {
			deficientsLog(i, ">", total, "=", SumExpr(is))
		}
	}
	fmt.Println("abundants", len(abundants))

	numbers := make([]bool, limit)
	for _, v1 := range abundants {
		for _, v2 := range abundants {
			if v1+v2 > limit {
				break
			}
			numbers[v1+v2-1] = true
		}
	}

	var sum, sumOfAbundants, nonSumOfAbundants int
	var countSumOfAbundants, countNonSumOfAbundants int
	fmt.Println("numbers", len(numbers))
	typesLog := fileLogln(ctx, "number-types.txt")
	for i, v := range numbers {
		number := i + 1
		sum += number
		if v {
			sumOfAbundants += number
			countSumOfAbundants++
			typesLog(number, "is sum of two abundants")
		} else {
			nonSumOfAbundants += number
			countNonSumOfAbundants++
			typesLog(number)
		}
	}
	fmt.Println("            Sums", sum)
	fmt.Println("No-Abundant Sums", nonSumOfAbundants, countNonSumOfAbundants)
	fmt.Println("   Abundant Sums", sumOfAbundants, countSumOfAbundants)
}

func has(ui uint, uis []uint) bool {
	for _, v := range uis {
		if v == ui {
			return true
		}
	}
	return false
}

func uints2Strs(is ...uint) []string {
	var ss []string = make([]string, len(is))
	for i, v := range is {
		ss[i] = fmt.Sprintf("%d", v)
	}
	return ss
}

type SumExpr []uint

func (se SumExpr) String() string {
	slices.Sort(se)
	var ss []string
	for _, i := range se {
		ss = append(ss, fmt.Sprintf("%d", i))
	}
	return strings.Join(ss, " + ")
}

func sum(is []uint) uint {
	var sum uint
	for _, i := range is {
		sum += i
	}
	return sum
}

func divisors(i uint) []uint {
	limit := uint(math.Sqrt(float64(i)))
	result := []uint{1}
	var j uint = 2
	for ; j <= limit; j++ {
		if i%uint(j) == 0 {
			result = append(result, uint(j))
			if v := i / j; v != 1 && v != j {
				result = append(result, v)
			}
		}
	}
	return result
}

type String []uint

func (s String) String() string {
	ss := make([]string, len(s))
	for i, v := range s {
		ss[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(ss, " + ")
}

func fileLogln(ctx context.Context, file string) func(anys ...any) {
	if !enableLog {
		return func(anys ...any) {}
	}
	f, err := os.OpenFile(file, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	go func() {
		<-ctx.Done()
		f.Close()
	}()
	return func(anys ...any) {
		fmt.Fprintln(f, anys...)
	}
}
