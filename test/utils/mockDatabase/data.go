package mockDatabase

import (
	"subLease/src/server/domain"
	"subLease/src/server/socialSecurityNumber"
	"time"
)

func GetSampleApartment() domain.Apartment {
	apartment := domain.CreateApartment(
		1104,
		"Norra Stationsgatan",
		117,
		"113 64",
		"Stockholm",
	)
	apartment.Id = 1

	return apartment
}

func GetSampleOwner(apartment domain.Apartment) domain.Owner {
	owner := domain.CreateOwner(
		"Simon",
		"Olofsson",
		socialSecurityNumber.Create(
			time.Date(1989, time.June, 1, 0, 0, 0, 0, time.Local),
			"071",
			1),
		[]domain.Apartment {apartment},
	)
	owner.Id = 1

	return owner
}

func GetSampleTenant() domain.Tenant {
	tenant := domain.Tenant{
		FirstName: "Simpa",
		LastName: "Lainen",
		SocialSecurityNumber: socialSecurityNumber.Create(
			time.Date(1989, time.June, 1, 0, 0, 0, 0, time.Local),
			"071",
			1,
		),
	}
	tenant.Id = 1

	return tenant
}

func GetSampleLeaseContract(owner domain.Owner, tenant domain.Tenant, apartment domain.Apartment) domain.LeaseContract  {
	leaseContract := domain.CreateLeaseContract(
		time.Date(2018, time.June, 1, 0, 0, 0, 0, time.Local),
		time.Date(2019, time.June, 1, 0, 0, 0, 0, time.Local),
		owner,
		tenant,
		apartment,
	)
	leaseContract.Id = 1

	return leaseContract
}