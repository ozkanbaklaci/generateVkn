package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println(generateVkn())

}
func random2(min, max int) int {
	return rand.Intn(max-min) + min
}

func exponent(x, y float64) float64 {
	return math.Pow(x, y)
}

func generateVkn() string {
	rand.Seed(time.Now().Unix())
	var (
		digits [10]int
		vkn    string
	)
	for i := 0; i < 9; i++ {
		digits[i] = random2(0, 10)
	}
	k := 9
	total := 0
	lastDigit := 0
	for i := 0; i < 9; i++ {
		x := (digits[i] + k) % 10
		var e int = int(exponent(2, float64(k)))
		y := (x * e) % 9
		if x != 0 && y == 0 {
			total += 9
		} else {
			total += y
		}
		k = k - 1
	}
	if (total % 10) == 0 {
		lastDigit = 0
	} else {
		lastDigit = (10 - (total % 10))
	}
	digits[9] = lastDigit
	for i := 0; i < len(digits); i++ {
		vkn += strconv.Itoa(digits[i])
	}
	return vkn
}
