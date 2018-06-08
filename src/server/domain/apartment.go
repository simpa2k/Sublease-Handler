package domain

import (
	"subLease/src/server/address"
)

type Apartment struct {
	Id 		int
	Number  int
	Address address.Address
}

func CreateApartment(number int, street string, streetNumber int, zipCode string, city string) Apartment {
	return Apartment {
		Number: number,
		Address: address.Create(street, streetNumber, zipCode, city),
	}
}

func (a *Apartment) Equal(other *Apartment) bool {
	return a.Id == other.Id && a.Number == other.Number && a.Address.Equal(&other.Address)
}