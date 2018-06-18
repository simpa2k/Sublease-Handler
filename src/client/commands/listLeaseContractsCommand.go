package commands

import (
	"fmt"
	"subLease/src/inMemoryDatabase"
)

type ListLeaseContracts struct {
	inMemoryDatabase inMemoryDatabase.InMemoryDatabase
}

func CreateListLeaseContracts(inMemoryData inMemoryDatabase.InMemoryDatabase) Command {
	return &ListLeaseContracts{
		inMemoryDatabase: inMemoryData,
	}
}

func (llc ListLeaseContracts) Stage() bool {
	for _, leaseContract := range llc.inMemoryDatabase.GetLeaseContracts() {
		fmt.Printf("%v\n", leaseContract)
	}

	return false
}

func (llc ListLeaseContracts) Execute() {}
func (llc ListLeaseContracts) Undo()    {}
