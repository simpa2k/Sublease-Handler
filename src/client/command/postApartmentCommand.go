package command

import (
	"subLease/src/inMemoryDatabase"
	"subLease/src/server/domain"
)

type PostApartment struct {
	inMemoryData     *inMemoryDatabase.InMemoryDatabase
	creatorFunction  func(inMemoryDatabase.InMemoryDatabase) domain.Apartment
	createdApartment *domain.Apartment
}

func CreatePostApartment(inMemoryData *inMemoryDatabase.InMemoryDatabase, creatorFunction func(inMemoryDatabase.InMemoryDatabase) domain.Apartment) Command {
	return &PostApartment{
		inMemoryData:    inMemoryData,
		creatorFunction: creatorFunction,
	}
}

func (plc *PostApartment) Stage() bool {
	if plc.createdApartment == nil {
		createdApartment := plc.creatorFunction(*plc.inMemoryData)
		createdApartmentId := plc.inMemoryData.CreateApartment(createdApartment)
		createdApartment.Id = createdApartmentId
		plc.createdApartment = &createdApartment
	} else {
		plc.inMemoryData.CreateApartment(*plc.createdApartment)
	}
	return true
}

func (plc *PostApartment) Execute() {

}

func (plc *PostApartment) Undo() {
	plc.inMemoryData.DeleteApartment(plc.createdApartment.Id)
}
