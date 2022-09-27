package tdd

import (
	"io"
	"net/http/httptest"
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
		name     string
		input    int
		expected string
	}{
		{
			name:     "When number in criteria == 1",
			input:    42,
			expected: "42 is an even number",
		},
		{
			name:     "When number in criteria != 1",
			input:    45,
			expected: "45 is an odd number",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, OddOrEven(test.input))
		})
	}
}

func TestCheckHealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		writer := httptest.NewRecorder()
		CheckHealth(writer, req)
		response := writer.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, "health check passed", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t,
			"text/plain; charset=utf-8",
			response.Header.Get("Content-Type"))
		assert.Equal(t, nil, err)
	})
}
