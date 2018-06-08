package server

import (
	"subLease/src/server/database"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"subLease/src/server/domain"
)

func getOwnersHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(database.GetOwners())
	}
}

func getOwnerHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		json.NewEncoder(w).Encode(database.GetOwner(id))
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
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		var owner domain.Owner
		_ = json.NewDecoder(r.Body).Decode(&owner)

		json.NewEncoder(w).Encode(database.UpdateOwner(id, owner))
	}
}

func deleteOwnerHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
