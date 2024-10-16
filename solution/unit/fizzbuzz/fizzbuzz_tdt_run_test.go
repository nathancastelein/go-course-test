package fizzbuzz

import (
	"testing"

	"github.com/nathancastelein/go-course-test/unit/fizzbuzz"
)

func TestFizzBuzzWithTableDrivenTestAndParallelism(t *testing.T) {
	// Arrange
	type test struct {
		input    int
		expected string
	}

	tests := map[string]*test{
		"should return same number when it is not multiple of three or five": {
			input:    1,
			expected: "1",
		},
		"should return Fizz when number is multiple of three": {
			input:    3,
			expected: "Fizz",
		},
		"should return Buzz when number is multiple of five": {
			input:    5,
			expected: "Buzz",
		},
		"should return FizzBuzz when number is multiple of three and five": {
			input:    30,
			expected: "FizzBuzz",
		},
	}

	// Act
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result := fizzbuzz.LongFizzbuzz(test.input)

			// Assert
			if result != test.expected {
				t.Fatalf("%s failed, expected %s, got %s", name, test.expected, result)
			}
		})
	}
}
