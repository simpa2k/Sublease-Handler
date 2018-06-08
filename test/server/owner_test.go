package server

import (
	"testing"
	"subLease/test/utils"
	"subLease/src/server/database"
	"encoding/json"
	"subLease/src/server/domain"
	"time"
	"bytes"
	"subLease/src/server/socialSecurityNumber"
	"subLease/test/utils/mockDatabase"
)

func TestGetOwners(t *testing.T) {
	utils.AssertRequestResponseMatchesOracle(t, "GET", "/owner", nil, func(db database.Database) ([]byte, error) {
		return json.Marshal(db.GetOwners())
	})
}

func TestGetOwner(t *testing.T) {
	utils.AssertRequestResponseMatchesOracle(t, "GET", "/owner/1", nil, func(db database.Database) ([]byte, error) {
		owner, _ := db.GetOwner(1)
		return json.Marshal(owner)
	})
}

func TestPostOwner(t *testing.T) {
	newOwner := domain.Owner{
		FirstName: "Sumon",
		LastName: "Olafsen",
		SocialSecurityNumber: socialSecurityNumber.Create(time.Date(1990, time.July, 2, 0, 0, 0, 0, time.Local), "017", 1),
		Apartments: []domain.Apartment {domain.CreateApartment(1101, "Norra Stationsgatan", 119,  "113 64", "Stockholm")},
	}

	jsonBytes, _ := json.Marshal(newOwner)

	r, db := utils.SetupServerWithMockDatabase()
	ownersBeforeRequest := db.GetOwners()

	res := utils.RequestToServer(r, "POST", "/owner", bytes.NewReader(jsonBytes))

	utils.AssertResponseMatchesOracle(t, res, func() ([]byte, error) {
		return json.Marshal(append(ownersBeforeRequest, newOwner))
	})

	if !utils.ContainsOwner(db.GetOwners(), newOwner) {
		t.Error("Lease contract was not saved.")
	}
}

func TestUpdateOwnerUpdatesAllValues(t *testing.T) {
	newFirstName := "Sumon"
	newLastName := "Olafsen"

	newSocialSecurityNumber := socialSecurityNumber.Create(time.Date(1990, time.July, 2, 0, 0, 0, 0, time.Local), "017", 1)
	newSocialSecurityNumberJSONBytes, _ := json.Marshal(newSocialSecurityNumber)

	newApartment := mockDatabase.GetSampleApartment()
	newApartment.Number = 1001
	newApartments := []domain.Apartment {newApartment}
	newApartmentsJSONBytes, _ := json.Marshal(newApartments)

	newOwner := domain.CreateOwner(
		newFirstName,
		newLastName,
		newSocialSecurityNumber,
		newApartments,
	)
	newOwner.Id = 1

	r, db := utils.SetupServerWithMockDatabase()
	ownerBeforeRequest, _ := db.GetOwner(1)

	completeUrl := utils.BuildQuery("/owner", []struct {Key string; Value string} {
		{"id", "1"},
		{"firstname", newFirstName},
		{"lastname", newLastName},
		{"socialsecuritynumber", string(newSocialSecurityNumberJSONBytes)},
		{"apartments", string(newApartmentsJSONBytes)},
	})

	res := utils.RequestToServer(r, "PUT", completeUrl, nil)

	utils.AssertResponseMatchesOracle(t, res, func() ([]byte, error) {
		return json.Marshal(newOwner)
	})

	currentOwner, _ := db.GetOwner(1)
	if !currentOwner.Equal(&newOwner) {
		t.Error("Owner was not updated")
	}

	if utils.ContainsOwner(db.GetOwners(), ownerBeforeRequest) {
		t.Error("Old owner was not removed.")
	}
}

func TestDeleteOwner(t *testing.T) {
	r, db := utils.SetupServerWithMockDatabase()
	owner, found := db.GetOwner(1)
	if !found {
		panic("Could not find the owner to delete.")
	}

	res := utils.RequestToServer(r, "DELETE", "/owner/1", nil)

	utils.AssertResponseMatchesOracle(t, res, func() ([]byte, error) {
		return json.Marshal(owner)
	})

	if utils.ContainsOwner(db.GetOwners(), owner) {
		t.Error("Lease contract was not removed.")
	}
}
