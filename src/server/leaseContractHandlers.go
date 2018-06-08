package server

import (
	"subLease/src/server/database"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
	"subLease/src/server/domain"
	"strings"
	"time"
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

func updateLeaseContractHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()
		id, _ := strconv.Atoi(queryValues.Get("id"))

		dateLayout := "2006-01-02 15:04:05.999999999 -0700 MST"
		from, err := time.Parse(dateLayout, queryValues.Get("from"))
		if err != nil {
			panic(err)
		}

		to, err := time.Parse(dateLayout, queryValues.Get("to"))
		if err != nil {
			panic(err)
		}

		var owner domain.Owner
		_ = json.NewDecoder(strings.NewReader(queryValues.Get("owner"))).Decode(&owner)
		var tenant domain.Tenant
		_ = json.NewDecoder(strings.NewReader(queryValues.Get("tenant"))).Decode(&tenant)
		var apartment domain.Apartment
		_ = json.NewDecoder(strings.NewReader(queryValues.Get("apartment"))).Decode(&apartment)

		leaseContractUpdate := domain.LeaseContractUpdate {
			From: &from,
			To: &to,
			Owner: &owner,
			Tenant: &tenant,
			Apartment: &apartment,
		}

		updatedLeaseContract, foundLeaseContractWithId := database.UpdateLeaseContract(id, leaseContractUpdate)
		if foundLeaseContractWithId {
			json.NewEncoder(w).Encode(updatedLeaseContract)
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

