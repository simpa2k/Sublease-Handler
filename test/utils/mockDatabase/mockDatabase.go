package mockDatabase

import (
	"subLease/src/server/domain"
	"subLease/src/server/database"
)

type mockDatabase struct {
	apartments []domain.Apartment
	owners []domain.Owner
	tenants []domain.Tenant
	leaseContracts []domain.LeaseContract
}

func Create() database.Database {
	myApartment := GetSampleApartment()
	me := GetSampleOwner(myApartment)
	tenant := GetSampleTenant()
	leaseContract := GetSampleLeaseContract( me, tenant, myApartment)

	return &mockDatabase {
		[]domain.Apartment {myApartment},
		[]domain.Owner {me},
		[]domain.Tenant {tenant},
		[]domain.LeaseContract {leaseContract},

	}
}
