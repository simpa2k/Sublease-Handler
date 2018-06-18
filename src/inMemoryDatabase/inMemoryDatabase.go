package inMemoryDatabase

import (
	"subLease/src/server/domain"
)

type InMemoryDatabase struct {
	apartments           []domain.Apartment
	apartmentCounter     int
	owners               []domain.Owner
	ownerCounter         int
	tenants              []domain.Tenant
	tenantCounter        int
	leaseContracts       []domain.LeaseContract
	leaseContractCounter int
}

func CreateEmpty() InMemoryDatabase {
	return InMemoryDatabase{}
}

func CreateWithData(apartments []domain.Apartment, owners []domain.Owner, tenants []domain.Tenant, leaseContracts []domain.LeaseContract) InMemoryDatabase {
	db := InMemoryDatabase{}

	for _, apartment := range apartments {
		db.CreateApartment(apartment)
	}
	for _, owner := range owners {
		db.CreateOwner(owner)
	}
	for _, tenant := range tenants {
		db.CreateTenant(tenant)
	}
	for _, leaseContract := range leaseContracts {
		db.CreateLeaseContract(leaseContract)
	}

	return db
}
