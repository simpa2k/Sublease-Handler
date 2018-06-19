package command

import (
	"subLease/src/inMemoryDatabase"
	"subLease/src/server/domain"
)

type PostOwner struct {
	inMemoryData    *inMemoryDatabase.InMemoryDatabase
	creatorFunction func(inMemoryDatabase.InMemoryDatabase) domain.Owner
	createdOwner    *domain.Owner
}

func CreatePostOwner(inMemoryData *inMemoryDatabase.InMemoryDatabase, creatorFunction func(inMemoryDatabase.InMemoryDatabase) domain.Owner) Command {
	return &PostOwner{
		inMemoryData:    inMemoryData,
		creatorFunction: creatorFunction,
	}
}

func (plc *PostOwner) Stage() bool {
	if plc.createdOwner == nil {
		createdOwner := plc.creatorFunction(*plc.inMemoryData)
		createdOwnerId := plc.inMemoryData.CreateOwner(createdOwner)
		createdOwner.Id = createdOwnerId
		plc.createdOwner = &createdOwner
	} else {
		plc.inMemoryData.CreateOwner(*plc.createdOwner)
	}
	return true
}

func (plc *PostOwner) Execute() {

}

func (plc *PostOwner) Undo() {
	plc.inMemoryData.DeleteOwner(plc.createdOwner.Id)
}
