package database

import (
	"time"
	"subLease/src/server/domain"
)

type LeaseContractUpdate struct {
	From      *time.Time
	To        *time.Time
	Owner     *int
	Tenant    *int
	Apartment *int
}

func (lcu *LeaseContractUpdate) UpdateLeaseContractWithValuesFrom(lc *domain.LeaseContract, database Database) {
	if lcu.From != nil {
		lc.From = *lcu.From
	}
	if lcu.To != nil {
		lc.To = *lcu.To
	}
	if lcu.Owner != nil {
		if owner, found := database.GetOwner(*lcu.Owner); found {
			lc.Owner = owner
		}
	}
	if lcu.Tenant != nil {
		if tenant, found := database.GetTenant(*lcu.Tenant); found {
			lc.Tenant = tenant
		}
	}
	if lcu.Apartment != nil {
		if apartment, found := database.GetApartment(*lcu.Apartment); found {
			lc.Apartment = apartment
		}
	}
}

