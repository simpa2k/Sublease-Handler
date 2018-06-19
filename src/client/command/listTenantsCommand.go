package command

import (
	"fmt"
	"subLease/src/inMemoryDatabase"
)

type ListTenants struct {
	inMemoryDatabase inMemoryDatabase.InMemoryDatabase
}

func CreateListTenants(inMemoryData inMemoryDatabase.InMemoryDatabase) Command {
	return &ListTenants{
		inMemoryDatabase: inMemoryData,
	}
}

func (llc ListTenants) Stage() bool {
	for _, tenant := range llc.inMemoryDatabase.GetTenants() {
		fmt.Printf("%v\n", tenant)
	}

	return false
}

func (llc ListTenants) Execute() {}
func (llc ListTenants) Undo()    {}
