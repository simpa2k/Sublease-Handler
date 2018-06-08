package server

import (
	"testing"
	"encoding/json"
	"subLease/test/utils"
	"subLease/src/server/database"
	"subLease/src/server/domain"
	"time"
	"subLease/test/utils/mockDatabase"
	"bytes"
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
	newLeaseContract := domain.LeaseContract {
		From: time.Date(2018, time.July, 15, 0, 0, 0, 0, time.Local),
		To: time.Date(2019, time.July, 15, 0, 0, 0, 0, time.Local),
		Owner: mockDatabase.GetSampleOwner(mockDatabase.GetSampleApartment()),
		Tenant: mockDatabase.GetSampleTenant(),
		Apartment: mockDatabase.GetSampleApartment(),
	}

	jsonBytes, _ := json.Marshal(newLeaseContract)

	r, db := utils.SetupServerWithMockDatabase()
	leaseContractsBeforeRequest := db.GetLeaseContracts()

	res := utils.RequestToServer(r, "POST", "/lease_contract", bytes.NewReader(jsonBytes))

	utils.AssertResponseMatchesOracle(t, res, func() ([]byte, error) {
		return json.Marshal(append(leaseContractsBeforeRequest, newLeaseContract))
	})

	if !utils.Contains(db.GetLeaseContracts(), newLeaseContract) {
		t.Error("Lease contract was not saved.")
	}
}

func TestUpdateLeaseContractUpdatesAllValues(t *testing.T) {
	newFrom := time.Date(2018, time.July, 16, 0, 0, 0, 0, time.Local)
	newTo := time.Date(2019, time.July, 16, 0, 0, 0, 0, time.Local)

	newOwner := mockDatabase.GetSampleOwner(mockDatabase.GetSampleApartment())
	newOwner.FirstName = "Sumon"
	newOwnerJSONBytes, _ := json.Marshal(newOwner)

	newTenant := mockDatabase.GetSampleTenant()
	newTenant.FirstName = "Slemon"
	newTenantJSONBytes, _ := json.Marshal(newTenant)

	newApartment := mockDatabase.GetSampleApartment()
	newApartment.Number = 1001
	newApartmentJSONBytes, _ := json.Marshal(newApartment)

	newLeaseContract := domain.CreateLeaseContract(
		newFrom,
		newTo,
		newOwner,
		newTenant,
		newApartment,
	)
	newLeaseContract.Id = 1

	r, db := utils.SetupServerWithMockDatabase()
	leaseContractBeforeRequest, _ := db.GetLeaseContract(1)

	completeUrl := utils.BuildQuery("/lease_contract", []struct {Key string; Value string} {
		{"id", "1"},
		{"from", newFrom.String()},
		{"to", newTo.String()},
		{"owner", string(newOwnerJSONBytes)},
		{"tenant", string(newTenantJSONBytes)},
		{"apartment", string(newApartmentJSONBytes)},
	})

	res := utils.RequestToServer(r, "PUT", completeUrl, nil)

	utils.AssertResponseMatchesOracle(t, res, func() ([]byte, error) {
		return json.Marshal(newLeaseContract)
	})

	currentLeaseContract, _ := db.GetLeaseContract(1)
	if !currentLeaseContract.Equal(&newLeaseContract) {
		t.Error("Lease contract was not updated")
	}

	if utils.Contains(db.GetLeaseContracts(), leaseContractBeforeRequest) {
		t.Error("Old lease contract was not removed.")
	}
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

	if utils.Contains(db.GetLeaseContracts(), leaseContractToBeRemoved) {
		t.Error("Lease contract was not removed.")
	}
}