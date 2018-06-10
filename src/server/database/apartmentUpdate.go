package database

import (
	"subLease/src/server/address"
	"subLease/src/server/domain"
)

type ApartmentUpdate struct {
	Id 		*int
	Number  *int
	Address *address.Address
}

func(au *ApartmentUpdate) UpdateApartmentWithValuesFrom(a *domain.Apartment, database Database) {

}
