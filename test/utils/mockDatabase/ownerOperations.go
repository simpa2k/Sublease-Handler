package mockDatabase

import "subLease/src/server/domain"

func (d mockDatabase) GetOwners() []domain.Owner {
	return d.owners
}

func (d mockDatabase) GetOwner(id int) domain.Owner {
	return domain.Owner{}
}

func (d mockDatabase) CreateOwner(owner domain.Owner) []domain.Owner {
	return d.owners
}

func (d mockDatabase) UpdateOwner(id int, newOwner domain.Owner) domain.Owner {
	return domain.Owner{}
}

func (d mockDatabase) DeleteOwner(id int) domain.Owner {
	return domain.Owner{}
}

