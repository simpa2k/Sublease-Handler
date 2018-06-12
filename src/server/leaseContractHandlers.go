// Generated by text/template; DO NOT EDIT
package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"subLease/src/server/database"
	"subLease/src/server/domain"
	"time"

	"github.com/gorilla/mux"
)

func getLeaseContractsHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(database.GetLeaseContracts())
	}
}

func getLeaseContractHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		if leaseContract, found := database.GetLeaseContract(id); found {
			json.NewEncoder(w).Encode(leaseContract)
		}
	}
}

func createLeaseContractHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var leaseContract domain.LeaseContract
		_ = json.NewDecoder(r.Body).Decode(&leaseContract)

		json.NewEncoder(w).Encode(database.CreateLeaseContract(leaseContract))
	}
}

func updateLeaseContractHandler(db database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int

		queryValues := r.URL.Query()
		leaseContractUpdate := database.LeaseContractUpdate{}

		retrieveInt("id", queryValues, func(idString string) (int, error) {
			return strconv.Atoi(idString)
		}, func(parsedId int) {
			id = parsedId
		})

		retrieveTime("from", queryValues, func(fromString string) (time.Time, error) {
			return time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", fromString)
		}, func(parsedFrom time.Time) {
			leaseContractUpdate.From = &parsedFrom
		})

		retrieveTime("to", queryValues, func(toString string) (time.Time, error) {
			return time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", toString)
		}, func(parsedTo time.Time) {
			leaseContractUpdate.To = &parsedTo
		})

		retrieveInt("owner", queryValues, func(ownerString string) (int, error) {
			return strconv.Atoi(ownerString)
		}, func(parsedOwner int) {
			leaseContractUpdate.Owner = &parsedOwner
		})

		retrieveInt("tenant", queryValues, func(tenantString string) (int, error) {
			return strconv.Atoi(tenantString)
		}, func(parsedTenant int) {
			leaseContractUpdate.Tenant = &parsedTenant
		})

		retrieveInt("apartment", queryValues, func(apartmentString string) (int, error) {
			return strconv.Atoi(apartmentString)
		}, func(parsedApartment int) {
			leaseContractUpdate.Apartment = &parsedApartment
		})

		updatedLeaseContract, foundLeaseContractWithId := db.UpdateLeaseContract(id, leaseContractUpdate)
		if foundLeaseContractWithId {
			json.NewEncoder(w).Encode(updatedLeaseContract)
		} else {
			http.Error(w, "No lease contract with that id was found.", http.StatusBadRequest)
		}
	}
}

func deleteLeaseContractHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		if leaseContract, found := database.DeleteLeaseContract(id); found {
			json.NewEncoder(w).Encode(leaseContract)
		}
	}
}
