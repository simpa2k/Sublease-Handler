package utils

import "subLease/src/server/domain"

func Contains(leaseContracts []domain.LeaseContract, leaseContract domain.LeaseContract) bool {
	for i := range leaseContracts {
		if leaseContracts[i].Equal(&leaseContract) {
			return true
		}
	}
	return false
}
