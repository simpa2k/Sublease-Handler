// Generated by text/template; DO NOT EDIT
package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"subLease/src/server/database"
	"subLease/src/server/domain"
	"subLease/src/server/socialSecurityNumber"

	"github.com/gorilla/mux"
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
		queryValues := r.URL.Query()
		id, _ := strconv.Atoi(queryValues.Get("id"))
		firstName := queryValues.Get("firstName")
		lastName := queryValues.Get("lastName")
		var socialSecurityNumber socialSecurityNumber.SocialSecurityNumber
		_ = json.NewDecoder(strings.NewReader(queryValues.Get("socialSecurityNumber"))).Decode(&socialSecurityNumber)

		tenantUpdate := database.TenantUpdate{
			FirstName:            &firstName,
			LastName:             &lastName,
			SocialSecurityNumber: &socialSecurityNumber,
		}

		updatedTenant, foundTenantWithId := db.UpdateTenant(id, tenantUpdate)
		if foundTenantWithId {
			json.NewEncoder(w).Encode(updatedTenant)
		}
	}
}

func deleteTenantHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		if tenant, found := database.DeleteTenant(id); found {
			json.NewEncoder(w).Encode(tenant)
		}
	}
}
