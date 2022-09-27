package tdd

import (
	"fmt"
	"math"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hello %s, Welcome!", name)
}

func OddOrEven(value int) string {
	criteria := math.Mod(float64(value), 2)
	if criteria == 1 {
		return fmt.Sprintf("%v is an odd number", value)
	}
	return fmt.Sprintf("%v is an even number", value)
}
