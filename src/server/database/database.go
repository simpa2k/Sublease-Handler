package database

import (
	"subLease/src/server/domain"
)

type Database interface {
	GetApartments() []domain.Apartment
	GetApartment(id int) (domain.Apartment, bool)
	CreateApartment(apartment domain.Apartment) []domain.Apartment
	UpdateApartment(id int, newApartment domain.Apartment) domain.Apartment
	DeleteApartment(id int) domain.Apartment

	GetOwners() []domain.Owner
	GetOwner(id int) (domain.Owner, bool)
	CreateOwner(owner domain.Owner) []domain.Owner
	UpdateOwner(id int, ownerUpdate domain.OwnerUpdate) (domain.Owner, bool)
	DeleteOwner(id int) (domain.Owner, bool)

	GetTenants() []domain.Tenant
	GetTenant(id int) (domain.Tenant, bool)
	CreateTenant(tenant domain.Tenant) []domain.Tenant
	UpdateTenant(id int, newTenant domain.Tenant) domain.Tenant
	DeleteTenant(id int) domain.Tenant

	GetLeaseContracts() []domain.LeaseContract
	GetLeaseContract(id int) (domain.LeaseContract, bool)
	CreateLeaseContract(leaseContract domain.LeaseContract) []domain.LeaseContract
	UpdateLeaseContract(id int, leaseContractUpdate LeaseContractUpdate) (domain.LeaseContract, bool)
	DeleteLeaseContract(id int) (domain.LeaseContract, bool)
}

type actualDatabase struct {}

func Create() Database {
	return &actualDatabase{}
}