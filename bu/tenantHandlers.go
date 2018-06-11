package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"subLease/src/server/database"
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
		if tenant, found := database.GetTenant(id); found {
			json.NewEncoder(w).Encode(tenant)
		}
	}
}

func createTenantHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var tenant domain.Tenant
		_ = json.NewDecoder(r.Body).Decode(&tenant)

		json.NewEncoder(w).Encode(database.CreateTenant(tenant))
	}
}

func updateTenantHandler(db database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		var tenant domain.Tenant
		_ = json.NewDecoder(r.Body).Decode(&tenant)

		updatedTenant, foundTenantWithId := db.UpdateTenant(id, database.TenantUpdate{})
		if foundTenantWithId {
			json.NewEncoder(w).Encode(updatedTenant)
		}
	}
}

func deleteTenantHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
