//go:generate go run ../../src/templates/main/generate.go ../../src/templates/ ../utils/mockDatabase/ ../../src/server/domain/ ../../src/server/database/ ../../src/server/

package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"subLease/src/server/database"
	"subLease/src/server/domain"
	"subLease/test/utils"
	"subLease/test/utils/mockDatabase"
	"testing"
	"time"
)

func TestGetLeaseContracts(t *testing.T) {
	utils.AssertRequestResponseMatchesOracle(t, "GET", "/lease_contract", nil, func(db database.Database) ([]byte, error) {
		return json.Marshal(db.GetLeaseContracts())
	})
}

func TestGetLeaseContract(t *testing.T) {
	utils.AssertRequestResponseMatchesOracle(t, "GET", "/lease_contract/1", nil, func(db database.Database) ([]byte, error) {
		leaseContract, _ := db.GetLeaseContract(1)
		return json.Marshal(leaseContract)
	})
}

func TestPostLeaseContract(t *testing.T) {
	newLeaseContract := domain.LeaseContract{
		From:      time.Date(2018, time.July, 15, 0, 0, 0, 0, time.Local),
		To:        time.Date(2019, time.July, 15, 0, 0, 0, 0, time.Local),
		Owner:     mockDatabase.GetSampleOwner1(mockDatabase.GetSampleApartment1()),
		Tenant:    mockDatabase.GetSampleTenant1(),
		Apartment: mockDatabase.GetSampleApartment1(),
	}

	jsonBytes, _ := json.Marshal(newLeaseContract)

	r, db := utils.SetupServerWithMockDatabase()
	leaseContractsBeforeRequest := db.GetLeaseContracts()

	res := utils.RequestToServer(r, "POST", "/lease_contract", bytes.NewReader(jsonBytes))

	utils.AssertResponseMatchesOracle(t, res, func() ([]byte, error) {
		return json.Marshal(append(leaseContractsBeforeRequest, newLeaseContract))
	})

	if !utils.ContainsLeaseContract(db.GetLeaseContracts(), newLeaseContract) {
		t.Error("Lease contract was not saved.")
	}
}

func TestUpdateLeaseContractUpdatesAllValues(t *testing.T) {
	newLeaseContract := getNewLeaseContract()
	completeUrl := utils.BuildQuery("/lease_contract", []struct {
		Key   string
		Value string
	}{
		{"id", "1"},
		{"from", newLeaseContract.From.String()},
		{"to", newLeaseContract.To.String()},
		{"owner", strconv.Itoa(newLeaseContract.Owner.Id)},
		{"tenant", strconv.Itoa(newLeaseContract.Tenant.Id)},
		{"apartment", strconv.Itoa(newLeaseContract.Apartment.Id)},
	})

	r, db := utils.SetupServerWithMockDatabase()
	leaseContractBeforeRequest, _ := db.GetLeaseContract(1)

	res := utils.RequestToServer(r, "PUT", completeUrl, nil)

	utils.AssertResponseMatchesOracle(t, res, func() ([]byte, error) {
		return json.Marshal(newLeaseContract)
	})

	currentLeaseContract, _ := db.GetLeaseContract(1)
	if !currentLeaseContract.Equal(&newLeaseContract) {
		t.Error("Lease contract was not updated.")
	}

	if utils.ContainsLeaseContract(db.GetLeaseContracts(), leaseContractBeforeRequest) {
		t.Error("Old lease contract was not removed.")
	}
}

func TestUpdateLeaseContractWithNoQueryParametersSetDoesNothing(t *testing.T) {
	completeUrl := utils.BuildQuery("/lease_contract", []struct {
		Key   string
		Value string
	}{
		{"id", "1"},
	})

	r, db := utils.SetupServerWithMockDatabase()
	leaseContractBeforeRequest, _ := db.GetLeaseContract(1)

	res := utils.RequestToServer(r, "PUT", completeUrl, nil)

	utils.AssertResponseMatchesOracle(t, res, func() ([]byte, error) {
		return json.Marshal(leaseContractBeforeRequest)
	})

	currentLeaseContract, _ := db.GetLeaseContract(1)
	if !currentLeaseContract.Equal(&leaseContractBeforeRequest) {
		t.Error("Lease contract was updated.")
	}
}

func TestAllUnparsableQueryParametersAreReportedAsErrors(t *testing.T) {
	completeUrl := utils.BuildQuery("/lease_contract", []struct {
		Key   string
		Value string
	}{
		{"id", "a"},
		{"from", "2/6 2018"},
		{"to", "2/6 2019"},
		{"owner", "NaN"},
		{"tenant", "y"},
		{"apartment", "&"},
	})

	r, db := utils.SetupServerWithMockDatabase()
	leaseContractBeforeRequest, _ := db.GetLeaseContract(1)

	res := utils.RequestToServer(r, "PUT", completeUrl, nil)

	utils.AssertResponseMatchesOracle(t, res, func() ([]byte, error) {
		return json.Marshal([]string{
			"Error while parsing query parameter id. Error: strconv.Atoi: parsing \"a\": invalid syntax",
			"Error while parsing query parameter from. Error: parsing time \"2/6 2018\" as \"2006-01-02 15:04:05.999999999 -0700 MST\": cannot parse \"2018\" as \"2006\"",
			"Error while parsing query parameter to. Error: parsing time \"2/6 2019\" as \"2006-01-02 15:04:05.999999999 -0700 MST\": cannot parse \"2019\" as \"2006\"",
			"Error while parsing query parameter owner. Error: strconv.Atoi: parsing \"NaN\": invalid syntax",
			"Error while parsing query parameter tenant. Error: strconv.Atoi: parsing \"y\": invalid syntax",
			"Error while parsing query parameter apartment. Error: strconv.Atoi: parsing \"\u0026\": invalid syntax",
		})
	})

	currentLeaseContract, _ := db.GetLeaseContract(1)
	if !currentLeaseContract.Equal(&leaseContractBeforeRequest) {
		t.Error("Lease contract was updated.")
	}
}

func TestUpdateNonExistingLeaseContractReturnsBadRequest(t *testing.T) {
	newLeaseContract := getNewLeaseContract()
	completeUrl := utils.BuildQuery("/lease_contract", []struct {
		Key   string
		Value string
	}{
		{"id", "2"},
		{"from", newLeaseContract.From.String()},
		{"to", newLeaseContract.To.String()},
		{"owner", strconv.Itoa(newLeaseContract.Owner.Id)},
		{"tenant", strconv.Itoa(newLeaseContract.Tenant.Id)},
		{"apartment", strconv.Itoa(newLeaseContract.Apartment.Id)},
	})

	r, _ := utils.SetupServerWithMockDatabase()

	res := utils.RequestToServer(r, "PUT", completeUrl, nil)
	if res.Code != http.StatusBadRequest {
		t.Error("Server did not respond with bad request.")
	}
}

func TestUpdateLeaseContractWithNonExistingOwnerReturnsBadRequest(t *testing.T) {
	completeUrl := utils.BuildQuery("/lease_contract", []struct {
		Key   string
		Value string
	}{
		{"id", "1"},
		{"owner", strconv.Itoa(3)},
	})

	r, db := utils.SetupServerWithMockDatabase()
	leaseContractBeforeRequest, _ := db.GetLeaseContract(1)

	res := utils.RequestToServer(r, "PUT", completeUrl, nil)
	if res.Code != http.StatusBadRequest {
		t.Error("Server did not respond with bad request.")
	}

	currentLeaseContract, _ := db.GetLeaseContract(1)
	if !currentLeaseContract.Equal(&leaseContractBeforeRequest) {
		t.Error("Lease contract was updated.")
	}
}

func TestUpdateLeaseContractWithNonExistingTenantReturnsBadRequest(t *testing.T) {
	completeUrl := utils.BuildQuery("/lease_contract", []struct {
		Key   string
		Value string
	}{
		{"id", "1"},
		{"tenant", strconv.Itoa(3)},
	})

	r, db := utils.SetupServerWithMockDatabase()
	leaseContractBeforeRequest, _ := db.GetLeaseContract(1)

	res := utils.RequestToServer(r, "PUT", completeUrl, nil)
	if res.Code != http.StatusBadRequest {
		t.Error("Server did not respond with bad request.")
	}

	currentLeaseContract, _ := db.GetLeaseContract(1)
	if !currentLeaseContract.Equal(&leaseContractBeforeRequest) {
		t.Error("Lease contract was updated.")
	}
}

func TestUpdateLeaseContractWithNonExistingApartmentReturnsBadRequest(t *testing.T) {
	completeUrl := utils.BuildQuery("/lease_contract", []struct {
		Key   string
		Value string
	}{
		{"id", "1"},
		{"apartment", strconv.Itoa(3)},
	})

	r, db := utils.SetupServerWithMockDatabase()
	leaseContractBeforeRequest, _ := db.GetLeaseContract(1)

	res := utils.RequestToServer(r, "PUT", completeUrl, nil)
	if res.Code != http.StatusBadRequest {
		t.Error("Server did not respond with bad request.")
	}

	currentLeaseContract, _ := db.GetLeaseContract(1)
	if !currentLeaseContract.Equal(&leaseContractBeforeRequest) {
		t.Error("Lease contract was updated.")
	}
}

func getNewLeaseContract() domain.LeaseContract {
	newFrom := time.Date(2018, time.July, 16, 0, 0, 0, 0, time.Local)
	newTo := time.Date(2019, time.July, 16, 0, 0, 0, 0, time.Local)
	newOwner := mockDatabase.GetSampleOwner2(mockDatabase.GetSampleApartment2())
	newTenant := mockDatabase.GetSampleTenant2()
	newApartment := mockDatabase.GetSampleApartment2()
	newLeaseContract := domain.CreateLeaseContract(
		newFrom,
		newTo,
		newOwner,
		newTenant,
		newApartment,
	)
	newLeaseContract.Id = 1

	return newLeaseContract
}

func TestDeleteLeaseContract(t *testing.T) {
	r, db := utils.SetupServerWithMockDatabase()
	leaseContractToBeRemoved, found := db.GetLeaseContract(1)
	if !found {
		panic("Could not find the lease contract to delete.")
	}

	res := utils.RequestToServer(r, "DELETE", "/lease_contract/1", nil)

	utils.AssertResponseMatchesOracle(t, res, func() ([]byte, error) {
		return json.Marshal(leaseContractToBeRemoved)
	})

	if utils.ContainsLeaseContract(db.GetLeaseContracts(), leaseContractToBeRemoved) {
		t.Error("Lease contract was not removed.")
	}
}
