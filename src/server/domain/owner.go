package domain

import (
	"subLease/src/server/socialSecurityNumber"
)

type Owner struct {
	Id 					 int
	FirstName 			 string
	LastName 			 string
	SocialSecurityNumber socialSecurityNumber.SocialSecurityNumber
	Apartments 			 []Apartment
}

type OwnerUpdate struct {
	FirstName 			 *string
	LastName 			 *string
	SocialSecurityNumber *socialSecurityNumber.SocialSecurityNumber
	Apartments 			 *[]Apartment
}

func CreateOwner(firstName string, lastName string, socialSecurityNumber socialSecurityNumber.SocialSecurityNumber, apartments []Apartment) Owner {
	return Owner {
			FirstName: firstName,
			LastName: lastName,
			SocialSecurityNumber: socialSecurityNumber,
			Apartments: apartments,
	}
}

func (o *Owner) Equal(other *Owner) bool {
	return o.SocialSecurityNumber.Equal(&other.SocialSecurityNumber)
}

func (o *Owner) UpdateOwnerWithValuesFrom(ownerUpdate OwnerUpdate) {
	if ownerUpdate.FirstName != nil {
		o.FirstName = *ownerUpdate.FirstName
	}
	if ownerUpdate.LastName != nil {
		o.LastName = *ownerUpdate.LastName
	}
	if ownerUpdate.SocialSecurityNumber != nil {
		o.SocialSecurityNumber = *ownerUpdate.SocialSecurityNumber
	}
	if ownerUpdate.Apartments != nil {
		o.Apartments = *ownerUpdate.Apartments
	}
}
