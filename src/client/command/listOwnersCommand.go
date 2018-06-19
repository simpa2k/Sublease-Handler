package command

import (
	"fmt"
	"subLease/src/inMemoryDatabase"
)

type ListOwners struct {
	inMemoryDatabase inMemoryDatabase.InMemoryDatabase
}

func CreateListOwners(inMemoryData inMemoryDatabase.InMemoryDatabase) Command {
	return &ListOwners{
		inMemoryDatabase: inMemoryData,
	}
}

func (llc ListOwners) Stage() bool {
	for _, owner := range llc.inMemoryDatabase.GetOwners() {
		fmt.Printf("%v\n", owner)
	}

	return false
}

func (llc ListOwners) Execute() {}
func (llc ListOwners) Undo()    {}
