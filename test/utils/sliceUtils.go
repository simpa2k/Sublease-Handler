package utils

import "subLease/src/server/domain"

func ContainsLeaseContract(leaseContracts []domain.LeaseContract, leaseContract domain.LeaseContract) bool {
	for i := range leaseContracts {
		if leaseContracts[i].Equal(&leaseContract) {
			return true
		}
	}
	return false
}

func ContainsOwner(owners []domain.Owner, owner domain.Owner) bool {
	for i := range owners {
		if owners[i].Equal(&owner) {
			return true
		}
	}
	return false
}
