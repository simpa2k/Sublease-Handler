package mockDatabase

import (
	"subLease/src/server/database"
	"subLease/src/server/domain"
)

type mockDatabase struct {
	apartments     []domain.Apartment
	owners         []domain.Owner
	tenants        []domain.Tenant
	leaseContracts []domain.LeaseContract
}

func Create() database.Database {
	apartment1 := GetSampleApartment1()
	apartment2 := GetSampleApartment2()

	owner1 := GetSampleOwner1(apartment1)
	owner2 := GetSampleOwner2(apartment2)

	tenant1 := GetSampleTenant1()
	tenant2 := GetSampleTenant2()

	leaseContract := GetSampleLeaseContract(owner1, tenant1, apartment1)

	return &mockDatabase{
		[]domain.Apartment{apartment1, GetSampleApartment2()},
		[]domain.Owner{owner1, owner2},
		[]domain.Tenant{tenant1, tenant2},
		[]domain.LeaseContract{leaseContract},
	}
}
