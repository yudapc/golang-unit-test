package tdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mocks
type mockAnimalCanRun struct {
	mock.Mock
}

func (m *mockAnimalCanRun) Run() string {
	args := m.Called()
	return args.String(0)
}

// End mocks

////////////////////////////////////////////////
// Test Implementation

func TestAnimalRunFast(t *testing.T) {
	t.Run("Animal can run fast", func(t *testing.T) {
		mockAnimal := new(mockAnimalCanRun)
		mockAnimal.On("Run").Return("Run faster")
		result := AnimalRunFast(mockAnimal)
		assert.Equal(t, "Run faster", result)
	})
	t.Run("Animal can not run", func(t *testing.T) {
		mockAnimal := new(mockAnimalCanRun)
		mockAnimal.On("Run").Return("Can not run")
		result := AnimalRunFast(mockAnimal)
		assert.Equal(t, "Can not run", result)
	})
}
