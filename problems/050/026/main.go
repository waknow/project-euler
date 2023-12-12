package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"
)

func main() {
	ss := []string{
		"0.5",
		"0.33333",
		"0.25",
		"0.2",
		"0.1666",
		"0.142857142857",
		"0.125",
		"0.11111",
		"0.1",
	}

	for _, s := range ss {
		r, found, err := recurringCycle(s)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(s, found, r)
	}

	fmt.Println(unitFractionDecimal(34, 20))

	start := time.Now()

	const maxPrecision = 50
	const maxDenominator = 50

	var maxLen int
	for i := 2; i <= maxDenominator; i++ {
		s := unitFractionDecimal(i, maxPrecision)
		idx := indexFirstNumber(s)
		if idx == -1 {
			continue
		}
		r, n, _ := recurringCycle(s[idx:])
		log.Printf("%2d: 1/%-3d = %s", n, i, s[:idx]+r)
		if n > maxLen {
			maxLen = n
		}
	}
	log.Println("elapsed:", time.Since(start))
}

func indexFirstNumber(s string) int {
	for i, r := range s {
		if r == '0' || r == '.' {
			continue
		}
		return i
	}
	return -1
}

func unitFractionDecimal(denominator, maxPrecision int) string {
	sb := strings.Builder{}
	sb.Write([]byte("0."))

	var left int = 1
	for i := 0; i < maxPrecision; i++ {
		n := int(math.Log10(float64(denominator))) - int(math.Log10(float64(left)))
		times := int(math.Pow10(n + 1))
		left = left * times

		quotient, remainder := left/denominator, left%denominator
		sb.Write([]byte(strings.Repeat("0", n-int(math.Log10(float64(quotient))))))
		sb.Write([]byte(fmt.Sprintf("%d", quotient)))

		// fmt.Printf("%d -> %2d = %5d / %d     %2d = %5d %% %d\n", i, quotient, left, denominator, remainder, left, denominator)
		if remainder == 0 {
			return sb.String()
		}
		left = remainder
	}
	return sb.String()
}

func recurringCycle(s string) (string, int, error) {
	if len(s) < 2 {
		return s, 0, fmt.Errorf("invalid input")
	}

	p, p1, offset := 1, 0, 0
	var matched bool
	for p < len(s) {
		p1, offset = 0, 0
		for {
			matched = s[p1+offset] == s[p+offset]

			if p1+1 == p || p1+offset+1 == p || p+offset+1 == len(s) {
				break
			}

			if matched {
				offset++
			} else {
				p1++
				offset = 0
			}
		}
		// log.Printf("len=%d p1=%d offset=%d p=%d", len(s), p1, offset, p)
		if matched && p1+offset+1 < p {
			matched = false
		}
		if matched {
			break
		}
		p++
	}

	if !matched {
		return s, 0, nil
	}

	// log.Printf("%s len(s)=%d p1=%d offset=%d p=%d", s, len(s), p1, offset, p)

	return fmt.Sprintf("%s(%s)", s[:p1], s[p1:p]), p - p1, nil
}
