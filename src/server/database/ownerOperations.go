// Generated by generate.go; do not edit manually
package database

import (
	"errors"
	"subLease/src/server/domain"
)

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

func (d actualDatabase) UpdateOwner(id int, ownerUpdate OwnerUpdate) (domain.Owner, error) {
	return domain.Owner{}, errors.New("not implemented")
}

func (d actualDatabase) DeleteOwner(id int) (domain.Owner, bool) {
	return domain.Owner{}, false
}
