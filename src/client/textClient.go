package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"subLease/src/client/command"
	"subLease/src/inMemoryDatabase"
	"subLease/src/server/address"
	"subLease/src/server/domain"
	"time"
	"subLease/src/server/socialSecurityNumber"
)

func Run(serverUrl string) {
	d := inMemoryDatabase.CreateWithData(
		getApartments(serverUrl),
		getOwners(serverUrl),
		getTenants(serverUrl),
		getLeaseContracts(serverUrl),
	)
	commandPipe := command.CommandPipe{}

	mainLoop(&d, &commandPipe)
}

func mainLoop(d *inMemoryDatabase.InMemoryDatabase, commandPipe *command.CommandPipe) {
	operations := map[string]func(){
		"las":  func() { commandPipe.Stage(command.CreateListApartments(*d)) },
		"ca":   func() { commandPipe.Stage(command.CreatePostApartment(d, createApartment)) },
		"llcs": func() { commandPipe.Stage(command.CreateListLeaseContracts(*d)) },
		"clc":  func() { commandPipe.Stage(command.CreatePostLeaseContract(d, createLeaseContract)) },
		"los": func() { commandPipe.Stage(command.CreateListOwners(*d)) },
		"co":  func() { commandPipe.Stage(command.CreatePostOwner(d, createOwner)) },
		"lts": func() { commandPipe.Stage(command.CreateListTenants(*d)) },
		"ct":  func() { commandPipe.Stage(command.CreatePostTenant(d, createTenant)) },
		"u":    commandPipe.Undo,
		"r":    commandPipe.Redo,
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		withoutNewline := strings.TrimSuffix(command, "\n")
		if operation, found := operations[withoutNewline]; found {
			operation()
		} else {
			fmt.Println("Invalid command")
		}
	}
}

func getApartments(url string) []domain.Apartment {
	body := get(url + "/apartment")

	defer body.Close()
	var apartment []domain.Apartment
	_ = json.NewDecoder(body).Decode(&apartment)

	return apartment
}

func getLeaseContracts(url string) []domain.LeaseContract {
	body := get(url + "/lease_contract")

	defer body.Close()
	var leaseContracts []domain.LeaseContract
	_ = json.NewDecoder(body).Decode(&leaseContracts)

	return leaseContracts
}

func getOwners(url string) []domain.Owner {
	body := get(url + "/owner")

	defer body.Close()
	var owners []domain.Owner
	_ = json.NewDecoder(body).Decode(&owners)

	return owners
}

func getTenants(url string) []domain.Tenant {
	body := get(url + "/tenant")

	defer body.Close()
	var tenants []domain.Tenant
	_ = json.NewDecoder(body).Decode(&tenants)

	return tenants
}

func get(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.Body
}

func createApartment(inMemoryDatabase.InMemoryDatabase) domain.Apartment {
	reader := bufio.NewReader(os.Stdin)
	apartmentNumber, err := strconv.Atoi(readValue(reader, "Apartment number"))
	handle(err)

	street := readValue(reader, "Street")

	streetNumber, err := strconv.Atoi(readValue(reader, "Street number"))
	handle(err)

	zipCode := readValue(reader, "Zip code")
	city := readValue(reader, "City")

	return domain.CreateApartment(apartmentNumber, address.Create(street, streetNumber, zipCode, city))
}

func createLeaseContract(database inMemoryDatabase.InMemoryDatabase) domain.LeaseContract {
	reader := bufio.NewReader(os.Stdin)
	from, err := time.Parse("2/1 2006", readValue(reader, "From"))
	handle(err)

	to, err := time.Parse("2/1 2006", readValue(reader, "To"))
	handle(err)

	ownerId, err := strconv.Atoi(readValue(reader, "Owner id"))
	handle(err)
	owner, _ := database.GetOwner(ownerId)

	tenantId, err := strconv.Atoi(readValue(reader, "Tenant id"))
	handle(err)
	tenant, _ := database.GetTenant(tenantId)

	apartmentId, err := strconv.Atoi(readValue(reader, "Apartment id"))
	handle(err)
	apartment, _ := database.GetApartment(apartmentId)

	return domain.CreateLeaseContract(from, to, owner, tenant, apartment)
}

func createOwner(database inMemoryDatabase.InMemoryDatabase) domain.Owner {
	reader := bufio.NewReader(os.Stdin)
	firstName, lastName, ssn := readNameAndSsn(reader)

	apartmentIdsString := readValue(reader, "Apartment ids")
	apartmentStringIds := strings.Split(apartmentIdsString, ", ")

	var apartmentIds []int
	for _, apartmentStringId := range apartmentStringIds {
		apartmentId, err := strconv.Atoi(apartmentStringId)
		handle(err)
		apartmentIds = append(apartmentIds, apartmentId)
	}

	apartments := database.GetApartmentsById(apartmentIds)
	if len(apartments) != len(apartmentIds) {
		panic("All apartment ids not found")
	}

	return domain.CreateOwner(firstName, lastName, ssn, apartments)
}

func createTenant(inMemoryDatabase.InMemoryDatabase) domain.Tenant {
	reader := bufio.NewReader(os.Stdin)
	firstName, lastName, ssn := readNameAndSsn(reader)

	return domain.CreateTenant(firstName, lastName, ssn)
}

func readNameAndSsn(reader *bufio.Reader) (string, string, socialSecurityNumber.SocialSecurityNumber) {
	firstName := readValue(reader, "First name")
	lastName := readValue(reader, "Last name")

	ssn, err := socialSecurityNumber.FromString(readValue(reader, "Social security number"))
	handle(err)

	return firstName, lastName, ssn
}

func readValue(reader *bufio.Reader, message string) string {
	fmt.Printf("%s > ", message)
	command, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(command, "\n")
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
