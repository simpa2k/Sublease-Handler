package server

import (
	"subLease/src/server/database"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"subLease/src/server/domain"
)

func getTenantsHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(database.GetTenants())
	}
}

func getTenantHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		json.NewEncoder(w).Encode(database.GetTenant(id))
	}
}

func createTenantHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var tenant domain.Tenant
		_ = json.NewDecoder(r.Body).Decode(&tenant)

		json.NewEncoder(w).Encode(database.CreateTenant(tenant))
	}
}

func updateTenantHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		var tenant domain.Tenant
		_ = json.NewDecoder(r.Body).Decode(&tenant)

		json.NewEncoder(w).Encode(database.UpdateTenant(id, tenant))
	}
}

func deleteTenantHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

