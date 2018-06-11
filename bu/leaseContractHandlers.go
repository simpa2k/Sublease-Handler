package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"subLease/src/server/database"
	"subLease/src/server/domain"
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

func updateLeaseContractHandler(db database.Database) func(w http.ResponseWriter, r *http.Request) {
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

		ownerId, err := strconv.Atoi(queryValues.Get("owner"))
		tenantId, err := strconv.Atoi(queryValues.Get("tenant"))
		apartmentId, err := strconv.Atoi(queryValues.Get("apartment"))

		leaseContractUpdate := database.LeaseContractUpdate{
			From:      &from,
			To:        &to,
			Owner:     &ownerId,
			Tenant:    &tenantId,
			Apartment: &apartmentId,
		}

		updatedLeaseContract, foundLeaseContractWithId := db.UpdateLeaseContract(id, leaseContractUpdate)
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
