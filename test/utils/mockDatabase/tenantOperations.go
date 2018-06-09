package mockDatabase

import "subLease/src/server/domain"

func (d mockDatabase) GetTenants() []domain.Tenant {
	return d.tenants
}

func findTenantById(tenants []domain.Tenant, id int) (*domain.Tenant, bool) {
	for _, tenant := range tenants {
		if tenant.Id == id {
			return &tenant, true
		}
	}
	return nil, false
}

func (d mockDatabase) GetTenant(id int) domain.Tenant {
	return domain.Tenant{}
}

func (d mockDatabase) CreateTenant(tenant domain.Tenant) []domain.Tenant {
	return d.tenants
}

func (d mockDatabase) UpdateTenant(id int, newTenant domain.Tenant) domain.Tenant {
	return domain.Tenant{}
}

func (d mockDatabase) DeleteTenant(id int) domain.Tenant {
	return domain.Tenant{}
}

