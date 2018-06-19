package command

import (
	"subLease/src/inMemoryDatabase"
	"subLease/src/server/domain"
)

type PostTenant struct {
	inMemoryData    *inMemoryDatabase.InMemoryDatabase
	creatorFunction func(inMemoryDatabase.InMemoryDatabase) domain.Tenant
	createdTenant   *domain.Tenant
}

func CreatePostTenant(inMemoryData *inMemoryDatabase.InMemoryDatabase, creatorFunction func(inMemoryDatabase.InMemoryDatabase) domain.Tenant) Command {
	return &PostTenant{
		inMemoryData:    inMemoryData,
		creatorFunction: creatorFunction,
	}
}

func (plc *PostTenant) Stage() bool {
	if plc.createdTenant == nil {
		createdTenant := plc.creatorFunction(*plc.inMemoryData)
		createdTenantId := plc.inMemoryData.CreateTenant(createdTenant)
		createdTenant.Id = createdTenantId
		plc.createdTenant = &createdTenant
	} else {
		plc.inMemoryData.CreateTenant(*plc.createdTenant)
	}
	return true
}

func (plc *PostTenant) Execute() {

}

func (plc *PostTenant) Undo() {
	plc.inMemoryData.DeleteTenant(plc.createdTenant.Id)
}
