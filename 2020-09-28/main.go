package sep282020

import (
	"fmt"
	"math"
)

type Person struct {
	name   string
	hunger uint
}

type People []Person

func Task() {
	peeps := People{
		newPerson("Joe", 9),
		newPerson("Cami", 3),
		newPerson("Cassidy", 4),
		newPerson("Gabor", 2),
	}

	slicesPerPizza := uint(7)

	pizzaz := gimmePizza(peeps, slicesPerPizza)
	fmt.Printf("28th September 2020\n-------------------\nWe need %d pizzas sliced into %d slices to satisfy %d people wanting %d total slices of pizza.\n\n", pizzaz, slicesPerPizza, len(peeps), peeps.TotalSlices())
}

// newPerson is a constructor for Person.
func newPerson(name string, hunger uint) Person {
	return Person{
		name:   name,
		hunger: hunger,
	}
}

// TotalSlices loops over the Persons in People to find out how many slices there are. It's to replace a map or reduce.
func (p People) TotalSlices() uint {
	total := uint(0)
	for _, person := range p {
		total += person.hunger
	}
	return total
}

// gimmePizza is the solution to the newsletter issue.
func gimmePizza(people People, pizzaCanBeSlicedInto uint) uint {
	return uint(math.Ceil(float64(people.TotalSlices()) / float64(pizzaCanBeSlicedInto)))
}
