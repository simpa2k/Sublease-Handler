package database

import "subLease/src/server/domain"

func (d actualDatabase) GetLeaseContracts() []domain.LeaseContract {
	return make([]domain.LeaseContract, 0, 0)
}

func (d actualDatabase) GetLeaseContract(id int) (domain.LeaseContract, bool) {
	return domain.LeaseContract{}, false
}

func (d actualDatabase) GetLeaseContractsById(ids []int) []domain.LeaseContract {
	return make([]domain.LeaseContract, 0, 0)
}

func (d actualDatabase) CreateLeaseContract(leaseContract domain.LeaseContract) []domain.LeaseContract {
	return make([]domain.LeaseContract, 0, 0)
}

func (d actualDatabase) UpdateLeaseContract(id int, leaseContractUpdate LeaseContractUpdate) (domain.LeaseContract, bool) {
	return domain.LeaseContract{}, false
}

func (d actualDatabase) DeleteLeaseContract(id int) (domain.LeaseContract, bool) {
	return domain.LeaseContract{}, false
}
