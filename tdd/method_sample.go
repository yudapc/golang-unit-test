package tdd

import "fmt"

type Animal interface {
	Run() string
}

// Class Duck
type duck struct {
	Name string
}

func NewDuck(name string) *duck {
	return &duck{
		Name: name,
	}
}

func (d *duck) Run() string {
	return fmt.Sprintf("Duck %s can not run", d.Name)
}

// End Class Duck

////////////////////////////////////////////////

// Class Dog
type dog struct {
	Name  string
	Color string
}

func NewDog(name, color string) *dog {
	return &dog{
		Name:  name,
		Color: color,
	}
}

func (dg *dog) Run() string {
	return fmt.Sprintf("Dog %s with color %s can run", dg.Name, dg.Color)
}

// End Class Dog

////////////////////////////////////////////////
// Implementation
////////////////////////////////////////////////

func AnimalRunFast(animal Animal) string {
	return animal.Run()
}

func DogRun() string {
	dog := NewDog("Chamar", "White")
	return AnimalRunFast(dog)
}

func DuckRun() string {
	duck := NewDuck("wek wek")
	return AnimalRunFast(duck)
}
