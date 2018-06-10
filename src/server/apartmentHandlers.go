package server

import (
	"subLease/src/server/database"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"subLease/src/server/domain"
	"strconv"
)

func getApartmentsHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(database.GetApartments())
	}
}

func getApartmentHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		if apartment, found := database.GetApartment(id); found {
			json.NewEncoder(w).Encode(apartment)
		}
	}
}

func createApartmentHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var apartment domain.Apartment
		_ = json.NewDecoder(r.Body).Decode(&apartment)

		json.NewEncoder(w).Encode(database.CreateApartment(apartment))
	}
}

func updateApartmentHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		var apartment domain.Apartment
		_ = json.NewDecoder(r.Body).Decode(&apartment)

		json.NewEncoder(w).Encode(database.UpdateApartment(id, apartment))
	}
}

func deleteApartmentHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
