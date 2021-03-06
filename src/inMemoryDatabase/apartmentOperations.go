// Generated by generate.go; do not edit manually
package inMemoryDatabase

import (
	"errors"
	"subLease/src/server/database"
	"subLease/src/server/domain"
)

func (d InMemoryDatabase) GetApartments() []domain.Apartment {
	return d.apartments
}

func (d InMemoryDatabase) GetApartment(id int) (domain.Apartment, bool) {
	if apartment, found := findApartmentById(d.apartments, id); found {
		return *apartment, found
	}
	return domain.Apartment{}, false
}

func (d InMemoryDatabase) GetApartmentsById(ids []int) []domain.Apartment {
	var foundApartments []domain.Apartment
	for _, id := range ids {
		if apartment, found := findApartmentById(d.apartments, id); found {
			foundApartments = append(foundApartments, *apartment)
		}
	}
	return foundApartments
}

func findApartmentById(apartments []domain.Apartment, id int) (*domain.Apartment, bool) {
	for _, apartment := range apartments {
		if apartment.Id == id {
			return &apartment, true
		}
	}
	return nil, false
}

func (d *InMemoryDatabase) CreateApartment(apartment domain.Apartment) int {
	d.apartmentCounter++
	apartment.Id = d.apartmentCounter
	d.apartments = append(d.apartments, apartment)
	return apartment.Id
}

func (d *InMemoryDatabase) UpdateApartment(id int, apartmentUpdate database.ApartmentUpdate) (domain.Apartment, error) {
	entityToReturn := domain.Apartment{}
	var errorToReturn error
	if i := indexOfApartment(d.apartments, id); i != -1 {
		updated, err := apartmentUpdate.UpdateApartmentWithValuesFrom(d.apartments[i], d)
		if err != nil {
			errorToReturn = err
		} else {
			d.apartments[i] = updated
			entityToReturn = updated
		}
	} else {
		errorToReturn = errors.New("no Apartment with that id was found")
	}
	return entityToReturn, errorToReturn
}

func indexOfApartment(apartments []domain.Apartment, id int) int {
	for i, apartment := range apartments {
		if apartment.Id == id {
			return i
		}
	}
	return -1
}

func (d *InMemoryDatabase) DeleteApartment(id int) (domain.Apartment, bool) {
	apartmentToRemove := domain.Apartment{}
	found := false
	j := 0
	for _, apartment := range d.apartments {
		if apartment.Id != id {
			d.apartments[j] = apartment
			j++
		} else {
			apartmentToRemove = apartment
			d.apartmentCounter--
			found = true
		}
	}
	d.apartments = d.apartments[:j]
	return apartmentToRemove, found
}
