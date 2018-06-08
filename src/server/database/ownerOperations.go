package database

import "subLease/src/server/domain"

func (d actualDatabase) GetOwners() []domain.Owner {
	return make([]domain.Owner, 0, 0)
}

func (d actualDatabase) GetOwner(id int) domain.Owner {
	return domain.Owner{}
}

func (d actualDatabase) CreateOwner(owner domain.Owner) []domain.Owner {
	return make([]domain.Owner, 0, 0)
}

func (d actualDatabase) UpdateOwner(id int, newOwner domain.Owner) domain.Owner {
	return domain.Owner{}
}

func (d actualDatabase) DeleteOwner(id int) domain.Owner {
	return domain.Owner{}
}

