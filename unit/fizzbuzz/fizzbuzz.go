package fizzbuzz

import (
	"fmt"
	"time"
)

var (
	fizz           = "Fizz"
	buzz           = "Buzz"
	fizzMultiplier = 3
	buzzMultiplier = 5
)

func Fizzbuzz(input int) string {
	fizzBuzzValue := ""
	if isMultipleOf(input, fizzMultiplier) {
		fizzBuzzValue += fizz
	}

	if isMultipleOf(input, buzzMultiplier) {
		fizzBuzzValue += buzz
	}

	if fizzBuzzValue != "" {
		return fizzBuzzValue
	}

	return fmt.Sprintf("%d", input)
}

func isMultipleOf(number, divider int) bool {
	return number%divider == 0
}

func LongFizzbuzz(input int) string {
	time.Sleep(500 * time.Millisecond)
	return Fizzbuzz(input)
}
