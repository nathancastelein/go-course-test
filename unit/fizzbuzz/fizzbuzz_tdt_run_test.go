package fizzbuzz

import "testing"

/*
Remove the first line t.SkipNow() in the below function.

Add parallelism by using t.Run and t.Parallel functions:
	- https://pkg.go.dev/testing#T.Run
	- https://pkg.go.dev/testing#T.Parallel

Example:
	func TestWithSubTests(t *testing.T) {
		t.Run("first test", func(t *testing.T) {
			t.Parallel()
			// do some stuff
		})

		t.Run("second test", func(t *testing.T) {
			t.Parallel()
			// do some stuff
		})
	}
*/

func TestFizzBuzzWithTableDrivenTestAndParallelism(t *testing.T) {
	t.SkipNow() // REMOVE ME

	// Arrange
	type test struct {
		input    int
		expected string
	}

	tests := map[string]test{
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
		result := LongFizzbuzz(test.input)

		// Assert
		if result != test.expected {
			t.Fatalf("%s failed, expected %s, got %s", name, test.expected, result)
		}
	}
}
