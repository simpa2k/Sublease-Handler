package database

import (
	"subLease/src/server/domain"
)

type Database interface {
	GetApartments() []domain.Apartment
	GetApartment(id int) domain.Apartment
	CreateApartment(apartment domain.Apartment) []domain.Apartment
	UpdateApartment(id int, newApartment domain.Apartment) domain.Apartment
	DeleteApartment(id int) domain.Apartment

	GetOwners() []domain.Owner
	GetOwner(id int) domain.Owner
	CreateOwner(owner domain.Owner) []domain.Owner
	UpdateOwner(id int, newOwner domain.Owner) domain.Owner
	DeleteOwner(id int) domain.Owner

	GetTenants() []domain.Tenant
	GetTenant(id int) domain.Tenant
	CreateTenant(tenant domain.Tenant) []domain.Tenant
	UpdateTenant(id int, newTenant domain.Tenant) domain.Tenant
	DeleteTenant(id int) domain.Tenant

	GetLeaseContracts() []domain.LeaseContract
	GetLeaseContract(id int) (domain.LeaseContract, bool)
	CreateLeaseContract(leaseContract domain.LeaseContract) []domain.LeaseContract
	UpdateLeaseContract(id int, leaseContractUpdate domain.LeaseContractUpdate) (domain.LeaseContract, bool)
	DeleteLeaseContract(id int) (domain.LeaseContract, bool)
}

type actualDatabase struct {}

func Create() Database {
	return &actualDatabase{}
}