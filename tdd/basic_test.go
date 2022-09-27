package tdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	greeting := SayHello("Yuda")
	assert.Equal(t, "Hello Yuda, Welcome!", greeting)
}
