package mockDatabase

import "subLease/src/server/domain"

func (d mockDatabase) GetOwners() []domain.Owner {
	return d.owners
}

func (d mockDatabase) GetOwner(id int) (domain.Owner, bool) {
	if owner, found := findOwnerById(d.owners, id); found {
		return *owner, found
	}
	return domain.Owner{}, false
}

func findOwnerById(owners []domain.Owner, id int) (*domain.Owner, bool) {
	for _, owner := range owners{
		if owner.Id == id {
			return &owner, true
		}
	}
	return nil, false
}

func (d* mockDatabase) CreateOwner(owner domain.Owner) []domain.Owner {
	d.owners = append(d.owners, owner)
	return d.owners
}

func (d *mockDatabase) UpdateOwner(id int, ownerUpdate domain.OwnerUpdate) (domain.Owner, bool) {
	if i := indexOfOwner(d.owners, id); i != -1 {
		d.owners[i].UpdateOwnerWithValuesFrom(ownerUpdate)
		return d.owners[i], true
	}
	return domain.Owner{}, false
}

func indexOfOwner(owners []domain.Owner, id int) (int) {
	for i, owner := range owners {
		if owner.Id == id {
			return i
		}
	}
	return -1
}

func (d *mockDatabase) DeleteOwner(id int) (domain.Owner, bool) {
	ownerToRemove := domain.Owner{}
	found := false
	j := 0
	for _, owner := range d.owners {
		if owner.Id != id {
			d.owners[j] = owner
			j++
		} else {
			ownerToRemove = owner
			found = true
		}
	}
	d.owners = d.owners[:j]
	return ownerToRemove, found
}

