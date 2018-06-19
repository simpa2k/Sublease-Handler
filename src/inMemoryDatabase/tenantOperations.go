// Generated by generate.go; do not edit manually
package inMemoryDatabase

import (
	"errors"
	"subLease/src/server/database"
	"subLease/src/server/domain"
)

func (d InMemoryDatabase) GetTenants() []domain.Tenant {
	return d.tenants
}

func (d InMemoryDatabase) GetTenant(id int) (domain.Tenant, bool) {
	if tenant, found := findTenantById(d.tenants, id); found {
		return *tenant, found
	}
	return domain.Tenant{}, false
}

func (d InMemoryDatabase) GetTenantsById(ids []int) []domain.Tenant {
	var foundTenants []domain.Tenant
	for _, id := range ids {
		if tenant, found := findTenantById(d.tenants, id); found {
			foundTenants = append(foundTenants, *tenant)
		}
	}
	return foundTenants
}

func findTenantById(tenants []domain.Tenant, id int) (*domain.Tenant, bool) {
	for _, tenant := range tenants {
		if tenant.Id == id {
			return &tenant, true
		}
	}
	return nil, false
}

func (d *InMemoryDatabase) CreateTenant(tenant domain.Tenant) int {
	d.tenantCounter++
	tenant.Id = d.tenantCounter
	d.tenants = append(d.tenants, tenant)
	return tenant.Id
}

func (d *InMemoryDatabase) UpdateTenant(id int, tenantUpdate database.TenantUpdate) (domain.Tenant, error) {
	entityToReturn := domain.Tenant{}
	var errorToReturn error
	if i := indexOfTenant(d.tenants, id); i != -1 {
		updated, err := tenantUpdate.UpdateTenantWithValuesFrom(d.tenants[i], d)
		if err != nil {
			errorToReturn = err
		} else {
			d.tenants[i] = updated
			entityToReturn = updated
		}
	} else {
		errorToReturn = errors.New("no Tenant with that id was found")
	}
	return entityToReturn, errorToReturn
}

func indexOfTenant(tenants []domain.Tenant, id int) int {
	for i, tenant := range tenants {
		if tenant.Id == id {
			return i
		}
	}
	return -1
}

func (d *InMemoryDatabase) DeleteTenant(id int) (domain.Tenant, bool) {
	tenantToRemove := domain.Tenant{}
	found := false
	j := 0
	for _, tenant := range d.tenants {
		if tenant.Id != id {
			d.tenants[j] = tenant
			j++
		} else {
			tenantToRemove = tenant
			d.tenantCounter--
			found = true
		}
	}
	d.tenants = d.tenants[:j]
	return tenantToRemove, found
}