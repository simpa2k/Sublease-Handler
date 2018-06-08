package server

import (
	"subLease/src/server/database"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"subLease/src/server/domain"
	"strings"
	"subLease/src/server/socialSecurityNumber"
)

func getOwnersHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(database.GetOwners())
	}
}

func getOwnerHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		if owner, found := database.GetOwner(id); found {
			json.NewEncoder(w).Encode(owner)
		}
	}
}

func createOwnerHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var owner domain.Owner
		_ = json.NewDecoder(r.Body).Decode(&owner)

		json.NewEncoder(w).Encode(database.CreateOwner(owner))
	}
}

func updateOwnerHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()
		id, _ := strconv.Atoi(queryValues.Get("id"))

		firstName := queryValues.Get("firstname")
		lastName := queryValues.Get("lastname")

		var ssn socialSecurityNumber.SocialSecurityNumber
		_ = json.NewDecoder(strings.NewReader(queryValues.Get("socialsecuritynumber"))).Decode(&ssn)
		var apartments []domain.Apartment
		_ = json.NewDecoder(strings.NewReader(queryValues.Get("apartments"))).Decode(&apartments)

		ownerUpdate := domain.OwnerUpdate{
			FirstName: &firstName,
			LastName: &lastName,
			SocialSecurityNumber: &ssn,
			Apartments: &apartments,
		}

		updatedOwner, foundOwnerWithId := database.UpdateOwner(id, ownerUpdate)
		if foundOwnerWithId {
			json.NewEncoder(w).Encode(updatedOwner)
		}
	}
}

func deleteOwnerHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		if owner, found := database.DeleteOwner(id); found {
			json.NewEncoder(w).Encode(owner)
		}
	}
}
