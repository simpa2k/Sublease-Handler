package database

import "subLease/src/server/domain"

func (d actualDatabase) GetOwners() []domain.Owner {
	return make([]domain.Owner, 0, 0)
}

func (d actualDatabase) GetOwner(id int) (domain.Owner, bool) {
	return domain.Owner{}, false
}

func (d actualDatabase) GetOwnersById(ids []int) []domain.Owner {
	return make([]domain.Owner, 0, 0)
}

func (d actualDatabase) CreateOwner(owner domain.Owner) []domain.Owner {
	return make([]domain.Owner, 0, 0)
}

func (d actualDatabase) UpdateOwner(id int, ownerUpdate OwnerUpdate) (domain.Owner, bool) {
	return domain.Owner{}, false
}

func (d actualDatabase) DeleteOwner(id int) (domain.Owner, bool) {
	return domain.Owner{}, false
}
