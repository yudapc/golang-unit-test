package tdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	greeting := SayHello("Yuda")
	assert.Equal(t, "Hello Yuda, Welcome!", greeting)
}

func TestOddOrEven(t *testing.T) {
	assert.Equal(t, "45 is an odd number", OddOrEven(45))
	assert.Equal(t, "42 is an even number", OddOrEven(42))
}

func TestOddOrEvenSubSet(t *testing.T) {
	t.Run("When number in criteria == 1", func(t *testing.T) {
		assert.Equal(t, "42 is an even number", OddOrEven(42))
	})
	t.Run("When number in criteria != 1", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", OddOrEven(45))
	})
}

func TestOddOrEvenTableTest(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{
			input:    42,
			expected: "42 is an even number",
		},
		{
			input:    45,
			expected: "45 is an odd number",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, OddOrEven(test.input))
	}
}
