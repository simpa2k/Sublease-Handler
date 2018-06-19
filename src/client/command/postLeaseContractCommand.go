package command

import (
	"subLease/src/inMemoryDatabase"
	"subLease/src/server/domain"
)

type PostLeaseContract struct {
	inMemoryData         *inMemoryDatabase.InMemoryDatabase
	creatorFunction      func(inMemoryDatabase.InMemoryDatabase) domain.LeaseContract
	createdLeaseContract *domain.LeaseContract
}

func CreatePostLeaseContract(inMemoryData *inMemoryDatabase.InMemoryDatabase, creatorFunction func(inMemoryDatabase.InMemoryDatabase) domain.LeaseContract) Command {
	return &PostLeaseContract{
		inMemoryData:    inMemoryData,
		creatorFunction: creatorFunction,
	}
}

func (plc *PostLeaseContract) Stage() bool {
	if plc.createdLeaseContract == nil {
		createdLeaseContract := plc.creatorFunction(*plc.inMemoryData)
		createdLeaseContractId := plc.inMemoryData.CreateLeaseContract(createdLeaseContract)
		createdLeaseContract.Id = createdLeaseContractId
		plc.createdLeaseContract = &createdLeaseContract
	} else {
		plc.inMemoryData.CreateLeaseContract(*plc.createdLeaseContract)
	}
	return true
}

func (plc *PostLeaseContract) Execute() {

}

func (plc *PostLeaseContract) Undo() {
	plc.inMemoryData.DeleteLeaseContract(plc.createdLeaseContract.Id)
}
