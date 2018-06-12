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
		var errors []string
		var id int

		queryValues := r.URL.Query()
		tenantUpdate := database.TenantUpdate{}

		retrieveInt("id", queryValues, func(idString string) (int, error) {
			return strconv.Atoi(idString)
		}, &errors, func(parsedId int) {
			id = parsedId
		})

		retrieveString("firstName", queryValues, func(firstNameString string) (string, error) {
			return firstNameString, nil
		}, &errors, func(parsedFirstName string) {
			tenantUpdate.FirstName = &parsedFirstName
		})

		retrieveString("lastName", queryValues, func(lastNameString string) (string, error) {
			return lastNameString, nil
		}, &errors, func(parsedLastName string) {
			tenantUpdate.LastName = &parsedLastName
		})

		retrieveSocialSecurityNumber("socialSecurityNumber", queryValues, func(socialSecurityNumberString string) (socialSecurityNumber.SocialSecurityNumber, error) {
			var socialSecurityNumber socialSecurityNumber.SocialSecurityNumber
			err := json.NewDecoder(strings.NewReader(socialSecurityNumberString)).Decode(&socialSecurityNumber)
			return socialSecurityNumber, err
		}, &errors, func(parsedSocialSecurityNumber socialSecurityNumber.SocialSecurityNumber) {
			tenantUpdate.SocialSecurityNumber = &parsedSocialSecurityNumber
		})

		if len(errors) > 0 {
			jsonError, _ := json.Marshal(errors)
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, string(jsonError), http.StatusBadRequest)
			return
		}

		updatedTenant, foundTenantWithId := db.UpdateTenant(id, tenantUpdate)
		if foundTenantWithId {
			json.NewEncoder(w).Encode(updatedTenant)
		} else {
			http.Error(w, "No tenant with that id was found.", http.StatusBadRequest)
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
