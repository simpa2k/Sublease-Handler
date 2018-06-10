package database

import "subLease/src/server/domain"

func (d actualDatabase) GetTenants() []domain.Tenant {
	return make([]domain.Tenant, 0, 0)
}

func (d actualDatabase) GetTenant(id int) (domain.Tenant, bool) {
	return domain.Tenant{}, false
}

func (d actualDatabase) CreateTenant(tenant domain.Tenant) []domain.Tenant {
	return make([]domain.Tenant, 0, 0)
}

func (d actualDatabase) UpdateTenant(id int, newTenant domain.Tenant) domain.Tenant {
	return domain.Tenant{}
}

func (d actualDatabase) DeleteTenant(id int) domain.Tenant {
	return domain.Tenant{}
}

