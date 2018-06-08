package domain

import (
	"subLease/src/server/socialSecurityNumber"
)

type Tenant struct {
	Id 					 int
	FirstName 			 string
	LastName 			 string
	SocialSecurityNumber socialSecurityNumber.SocialSecurityNumber
}

func (t *Tenant) Equal(other *Tenant) bool {
	return t.SocialSecurityNumber.Equal(&other.SocialSecurityNumber)
}