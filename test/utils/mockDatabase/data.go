package mockDatabase

import (
	"subLease/src/server/address"
	"subLease/src/server/domain"
	"subLease/src/server/socialSecurityNumber"
	"time"
)

func GetSampleApartment1() domain.Apartment {
	apartment := domain.CreateApartment(
		1104,
		address.Create(
			"Norra Stationsgatan",
			117,
			"113 64",
			"Stockholm",
		),
	)
	apartment.Id = 1

	return apartment
}

func GetSampleApartment2() domain.Apartment {
	apartment := domain.CreateApartment(
		1103,
		address.Create(
			"Norra Stationsgatan",
			119,
			"113 64",
			"Stockholm",
		),
	)
	apartment.Id = 2

	return apartment
}

func GetSampleOwner1(apartment domain.Apartment) domain.Owner {
	owner := domain.CreateOwner(
		"Simon",
		"Olofsson",
		socialSecurityNumber.Create(
			time.Date(1989, time.June, 1, 0, 0, 0, 0, time.Local),
			"111",
			1),
		[]domain.Apartment{apartment},
	)
	owner.Id = 1

	return owner
}

func GetSampleOwner2(apartment domain.Apartment) domain.Owner {
	owner := domain.CreateOwner(
		"Sumon",
		"Olafsen",
		socialSecurityNumber.Create(
			time.Date(1990, time.July, 2, 0, 0, 0, 0, time.Local),
			"111",
			1),
		[]domain.Apartment{apartment},
	)
	owner.Id = 2

	return owner
}

func GetSampleTenant1() domain.Tenant {
	tenant := domain.Tenant{
		FirstName: "Simpa",
		LastName:  "Lainen",
		SocialSecurityNumber: socialSecurityNumber.Create(
			time.Date(1989, time.June, 1, 0, 0, 0, 0, time.Local),
			"111",
			1,
		),
	}
	tenant.Id = 1

	return tenant
}

func GetSampleTenant2() domain.Tenant {
	tenant := domain.Tenant{
		FirstName: "Slemon	",
		LastName: "Slemsson",
		SocialSecurityNumber: socialSecurityNumber.Create(
			time.Date(1066, time.October, 14, 0, 0, 0, 0, time.Local),
			"111",
			1,
		),
	}
	tenant.Id = 2

	return tenant
}

func GetSampleLeaseContract(owner domain.Owner, tenant domain.Tenant, apartment domain.Apartment) domain.LeaseContract {
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
