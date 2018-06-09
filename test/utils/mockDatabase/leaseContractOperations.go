package mockDatabase

import "subLease/src/server/domain"

func (d mockDatabase) GetLeaseContracts() []domain.LeaseContract {
	return d.leaseContracts
}

func (d mockDatabase) GetLeaseContract(id int) (domain.LeaseContract, bool) {
	if leaseContract, found := findLeaseContractById(d.leaseContracts, id); found {
		return *leaseContract, found
	}
	return domain.LeaseContract{}, false
}

func findLeaseContractById(leaseContracts []domain.LeaseContract, id int) (*domain.LeaseContract, bool) {
	for _, leaseContract := range leaseContracts {
		if leaseContract.Id == id {
			return &leaseContract, true
		}
	}
	return nil, false
}

func (d *mockDatabase) CreateLeaseContract(leaseContract domain.LeaseContract) []domain.LeaseContract {
	d.leaseContracts = append(d.leaseContracts, leaseContract)
	return d.leaseContracts
}

func (d *mockDatabase) UpdateLeaseContract(id int, leaseContractUpdate domain.LeaseContractUpdate) (domain.LeaseContract, bool) {
	if i := indexOfLeaseContract(d.leaseContracts, id); i != -1 {
		d.UpdateLeaseContractWithValuesFrom(&d.leaseContracts[i], leaseContractUpdate)
		return d.leaseContracts[i], true
	}
	return domain.LeaseContract{}, false
}

func indexOfLeaseContract(leaseContracts []domain.LeaseContract, id int) (int) {
	for i, leaseContract := range leaseContracts {
		if leaseContract.Id == id {
			return i
		}
	}
	return -1
}

func (d *mockDatabase) UpdateLeaseContractWithValuesFrom(lc *domain.LeaseContract, leaseContractUpdate domain.LeaseContractUpdate) {
	if leaseContractUpdate.From != nil {
		lc.From = *leaseContractUpdate.From
	}
	if leaseContractUpdate.To != nil {
		lc.To = *leaseContractUpdate.To
	}
	if leaseContractUpdate.Owner != nil {
		if owner, found := findOwnerById(d.owners, *leaseContractUpdate.Owner); found {
			lc.Owner = *owner
		}
	}
	if leaseContractUpdate.Tenant != nil {
		if tenant, found := findTenantById(d.tenants, *leaseContractUpdate.Tenant); found {
			lc.Tenant = *tenant
		}
	}
	if leaseContractUpdate.Apartment != nil {
		if apartment, found := findApartmentById(d.apartments, *leaseContractUpdate.Apartment); found {
			lc.Apartment = *apartment
		}
	}
}

func (d* mockDatabase) DeleteLeaseContract(id int) (domain.LeaseContract, bool) {
	leaseContractToRemove := domain.LeaseContract{}
	found := false
	j := 0
	for _, leaseContract := range d.leaseContracts {
		if leaseContract.Id != id {
			d.leaseContracts[j] = leaseContract
			j++
		} else {
			leaseContractToRemove = leaseContract
			found = true
		}
	}
	d.leaseContracts = d.leaseContracts[:j]
	return leaseContractToRemove, found
}

