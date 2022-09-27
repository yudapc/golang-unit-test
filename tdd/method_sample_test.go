package tdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Test Duck
func TestDuckRunMethod(t *testing.T) {
	duck := NewDuck("bebek")
	result := duck.Run()
	assert.Equal(t, "Duck bebek can not run", result)
}

// End Test Duck
////////////////////////////////////////////////
// Test Dog
func TestDogRunMethod(t *testing.T) {
	dg := NewDog("Gahar", "Yellow")
	result := dg.Run()
	assert.Equal(t, "Dog Gahar with color Yellow can run", result)
}

// End Test Dog
////////////////////////////////////////////////

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
		expected := "Run faster"
		mockAnimal.On("Run").Return(expected)
		result := AnimalRunFast(mockAnimal)
		assert.Equal(t, expected, result)
	})
	t.Run("Animal can not run", func(t *testing.T) {
		mockAnimal := new(mockAnimalCanRun)
		expected := "Can not run"
		mockAnimal.On("Run").Return(expected)
		result := AnimalRunFast(mockAnimal)
		assert.Equal(t, expected, result)
	})
}
