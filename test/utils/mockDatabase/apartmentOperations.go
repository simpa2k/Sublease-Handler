package mockDatabase

import (
	"subLease/src/server/domain"
	"subLease/src/server/database"
)

func (d mockDatabase) GetApartments() []domain.Apartment {
	return d.apartments
}

func (d mockDatabase) GetApartment(id int) (domain.Apartment, bool) {
	if apartment, found := findApartmentById(d.apartments, id); found {
		return *apartment, found
	}
	return domain.Apartment{}, false
}

func findApartmentById(apartments []domain.Apartment, id int) (*domain.Apartment, bool) {
	for _, apartment := range apartments {
		if apartment.Id == id {
			return &apartment, true
		}
	}
	return nil, false
}

func (d *mockDatabase) CreateApartment(apartment domain.Apartment) []domain.Apartment {
	d.apartments = append(d.apartments, apartment)
	return d.apartments
}

func (d *mockDatabase) UpdateApartment(id int, apartmentUpdate database.ApartmentUpdate) (domain.Apartment, bool) {
	if i := indexOfApartment(d.apartments, id); i != -1 {
		apartmentUpdate.UpdateApartmentWithValuesFrom(&d.apartments[i], d)
		return d.apartments[i], true
	}
	return domain.Apartment{}, false
}

func indexOfApartment(apartments []domain.Apartment, id int) (int) {
	for i, apartment := range apartments {
		if apartment.Id == id {
			return i
		}
	}
	return -1
}

func (d* mockDatabase) DeleteApartment(id int) (domain.Apartment, bool) {
	apartmentToRemove := domain.Apartment{}
	found := false
	j := 0
	for _, apartment := range d.apartments {
		if apartment.Id != id {
			d.apartments[j] = apartment
			j++
		} else {
			apartmentToRemove = apartment
			found = true
		}
	}
	d.apartments = d.apartments[:j]
	return apartmentToRemove, found
}

