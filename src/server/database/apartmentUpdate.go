// Generated by generate.go; do not edit manually
package database

import (
	"subLease/src/server/address"
	"subLease/src/server/domain"
)

type ApartmentUpdate struct {
	Id      *int
	Number  *int
	Address *address.Address
}

func (a *ApartmentUpdate) UpdateApartmentWithValuesFrom(e domain.Apartment, database Database) (domain.Apartment, error) {
	if a.Id != nil {
		e.Id = *a.Id
	}
	if a.Number != nil {
		e.Number = *a.Number
	}
	if a.Address != nil {
		e.Address = *a.Address
	}
	return e, nil
}
