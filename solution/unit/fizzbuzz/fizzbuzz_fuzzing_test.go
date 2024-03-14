package fizzbuzz

import (
	"strconv"
	"testing"

	"github.com/nathancastelein/go-course-test/unit/fizzbuzz"
)

func FuzzFizzbuzz(f *testing.F) {
	f.Add(1) // Use f.Add to provide a seed corpus
	f.Add(3)
	f.Add(5)
	f.Add(30)

	f.Fuzz(func(t *testing.T, number int) {
		result := fizzbuzz.Fizzbuzz(number)
		switch result {
		case "Fizz", "Buzz", "FizzBuzz":
			return
		default:
			numberFromString, err := strconv.Atoi(result)
			if err != nil {
				t.Fatalf("got invalid result for input %d: %s (error %s)", number, result, err)
			}

			if numberFromString != number {
				t.Fatalf("number should not be different: expected %d got %d", number, numberFromString)
			}
		}
	})
}
