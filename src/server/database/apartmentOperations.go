package database

import "subLease/src/server/domain"

func (d actualDatabase) GetApartments() []domain.Apartment {
	return make([]domain.Apartment, 0, 0)
}

func (d actualDatabase) GetApartment(id int) (domain.Apartment, bool) {
	return domain.Apartment{}, false
}

func (d actualDatabase) GetApartmentsById(ids []int) []domain.Apartment {
	return make([]domain.Apartment, 0, 0)
}

func (d actualDatabase) CreateApartment(apartment domain.Apartment) []domain.Apartment {
	return make([]domain.Apartment, 0, 0)
}

func (d actualDatabase) UpdateApartment(id int, apartmentUpdate ApartmentUpdate) (domain.Apartment, bool) {
	return domain.Apartment{}, false
}

func (d actualDatabase) DeleteApartment(id int) (domain.Apartment, bool) {
	return domain.Apartment{}, false
}
