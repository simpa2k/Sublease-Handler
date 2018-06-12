// Generated by text/template; DO NOT EDIT
package domain




type LeaseContract struct {
    Id int
    From time.Time
    To time.Time
    Owner Owner
    Tenant Tenant
    Apartment Apartment
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
    return lc.Id == other.Id && lc.From == other.From && lc.To == other.To && lc.Owner.Equal(&other.Owner) && lc.Tenant.Equal(&other.Tenant) && lc.Apartment.Equal(&other.Apartment)
}