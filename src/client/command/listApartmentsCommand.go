package command

import (
	"fmt"
	"subLease/src/inMemoryDatabase"
)

type ListApartments struct {
	inMemoryDatabase inMemoryDatabase.InMemoryDatabase
}

func CreateListApartments(inMemoryData inMemoryDatabase.InMemoryDatabase) Command {
	return &ListApartments{
		inMemoryDatabase: inMemoryData,
	}
}

func (llc ListApartments) Stage() bool {
	for _, apartment := range llc.inMemoryDatabase.GetApartments() {
		fmt.Printf("%v\n", apartment)
	}

	return false
}

func (llc ListApartments) Execute() {}
func (llc ListApartments) Undo()    {}
