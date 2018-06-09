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
	Owner     *int
	Tenant    *int
	Apartment *int
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
