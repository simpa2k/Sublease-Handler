package mockDatabase

import "subLease/src/server/domain"

func (d mockDatabase) GetApartments() []domain.Apartment {
	return d.apartments
}

func (d mockDatabase) GetApartment(id int) domain.Apartment {
	return domain.Apartment{}
}

func (d mockDatabase) CreateApartment(apartment domain.Apartment) []domain.Apartment {
	return d.apartments
}

func (d mockDatabase) UpdateApartment(id int, newApartment domain.Apartment) domain.Apartment {
	return domain.Apartment{}
}

func (d mockDatabase) DeleteApartment(id int) domain.Apartment {
	return domain.Apartment{}
}

