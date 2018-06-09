package mockDatabase

import "subLease/src/server/domain"

func (d mockDatabase) GetApartments() []domain.Apartment {
	return d.apartments
}

func findApartmentById(apartments []domain.Apartment, id int) (*domain.Apartment, bool) {
	for _, apartment := range apartments {
		if apartment.Id == id {
			return &apartment, true
		}
	}
	return nil, false
}

func findApartmentsById(allApartments []domain.Apartment, ids []int) ([]domain.Apartment) {
	var foundApartments []domain.Apartment
	for _, id := range ids {
		if apartment, found := findApartmentById(allApartments, id); found {
			foundApartments = append(foundApartments, *apartment)
		}
	}
	return foundApartments
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

