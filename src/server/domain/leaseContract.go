package domain

import "time"

type LeaseContract struct {
	Id 		  int
	From      time.Time
	To        time.Time
	Owner     Owner
	Tenant    Tenant
	Apartment Apartment
}

type LeaseContractUpdate struct {
	From      *time.Time
	To        *time.Time
	Owner     *Owner
	Tenant    *Tenant
	Apartment *Apartment
}

func CreateLeaseContract(from time.Time, to time.Time, owner Owner, tenant Tenant, apartment Apartment) LeaseContract {
	return LeaseContract {
		From: from,
		To: to,
		Owner: owner,
		Tenant: tenant,
		Apartment: apartment,
	}
}

func (lc *LeaseContract) Equal(other *LeaseContract) bool {
	return lc.Id == other.Id &&
		lc.From.Equal(other.From) &&
		lc.To.Equal(other.To) &&
		lc.Owner.Equal(&other.Owner) &&
		lc.Tenant.Equal(&other.Tenant) &&
		lc.Apartment.Equal(&other.Apartment)
}

func (lc *LeaseContract) UpdateLeastContractWithValuesFrom(leaseContractUpdate LeaseContractUpdate) {
	if leaseContractUpdate.From != nil {
		lc.From = *leaseContractUpdate.From
	}
	if leaseContractUpdate.To != nil {
		lc.To = *leaseContractUpdate.To
	}
	if leaseContractUpdate.Owner != nil {
		lc.Owner = *leaseContractUpdate.Owner
	}
	if leaseContractUpdate.Tenant != nil {
		lc.Tenant = *leaseContractUpdate.Tenant
	}
	if leaseContractUpdate.Apartment != nil {
		lc.Apartment = *leaseContractUpdate.Apartment
	}
}

