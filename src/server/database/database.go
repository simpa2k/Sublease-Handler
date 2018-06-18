// Generated by generate.go; do not edit manually
package database

import (
	"subLease/src/server/domain"
)

type Database interface {
	GetApartments() []domain.Apartment
	GetApartment(id int) (domain.Apartment, bool)
	GetApartmentsById(id []int) []domain.Apartment
	CreateApartment(apartment domain.Apartment) int
	UpdateApartment(id int, apartmentUpdate ApartmentUpdate) (domain.Apartment, error)
	DeleteApartment(id int) (domain.Apartment, bool)

	GetLeaseContracts() []domain.LeaseContract
	GetLeaseContract(id int) (domain.LeaseContract, bool)
	GetLeaseContractsById(id []int) []domain.LeaseContract
	CreateLeaseContract(leaseContract domain.LeaseContract) int
	UpdateLeaseContract(id int, leaseContractUpdate LeaseContractUpdate) (domain.LeaseContract, error)
	DeleteLeaseContract(id int) (domain.LeaseContract, bool)

	GetOwners() []domain.Owner
	GetOwner(id int) (domain.Owner, bool)
	GetOwnersById(id []int) []domain.Owner
	CreateOwner(owner domain.Owner) int
	UpdateOwner(id int, ownerUpdate OwnerUpdate) (domain.Owner, error)
	DeleteOwner(id int) (domain.Owner, bool)

	GetTenants() []domain.Tenant
	GetTenant(id int) (domain.Tenant, bool)
	GetTenantsById(id []int) []domain.Tenant
	CreateTenant(tenant domain.Tenant) int
	UpdateTenant(id int, tenantUpdate TenantUpdate) (domain.Tenant, error)
	DeleteTenant(id int) (domain.Tenant, bool)
}

type actualDatabase struct{}

func Create() Database {
	return &actualDatabase{}
}
