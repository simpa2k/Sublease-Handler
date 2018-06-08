package database

import "subLease/src/server/domain"

func (d actualDatabase) GetApartments() []domain.Apartment {
	return make([]domain.Apartment, 0, 0)
}

func (d actualDatabase) GetApartment(id int) domain.Apartment {
	return domain.Apartment{}
}

func (d actualDatabase) CreateApartment(apartment domain.Apartment) []domain.Apartment {
	return make([]domain.Apartment, 0, 0)
}

func (d actualDatabase) UpdateApartment(id int, newApartment domain.Apartment) domain.Apartment {
	return domain.Apartment{}
}

func (d actualDatabase) DeleteApartment(id int) domain.Apartment {
	return domain.Apartment{}
}

