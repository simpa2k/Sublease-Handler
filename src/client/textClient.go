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
	"subLease/src/client/commands"
	"subLease/src/inMemoryDatabase"
	"subLease/src/server/domain"
	"time"
)

func Run(serverUrl string) {
	d := inMemoryDatabase.CreateWithData(
		getApartments(serverUrl),
		getOwners(serverUrl),
		getTenants(serverUrl),
		getLeaseContracts(serverUrl),
	)
	commandPipe := commands.CommandPipe{}

	mainLoop(&d, &commandPipe)
}

func mainLoop(d *inMemoryDatabase.InMemoryDatabase, commandPipe *commands.CommandPipe) {
	operations := map[string]func(){
		"llcs": func() { commandPipe.Stage(commands.CreateListLeaseContracts(*d)) },
		"clc":  func() { commandPipe.Stage(commands.CreatePostLeaseContract(d, createLeaseContract)) },
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

func createLeaseContract(database inMemoryDatabase.InMemoryDatabase) domain.LeaseContract {
	reader := bufio.NewReader(os.Stdin)
	from, err := time.Parse("2/1 2006", readValue(reader, "From"))
	if err != nil {
		panic(err)
	}

	to, err := time.Parse("2/1 2006", readValue(reader, "To"))
	if err != nil {
		panic(err)
	}

	ownerId, err := strconv.Atoi(readValue(reader, "Owner id"))
	if err != nil {
		panic(err)
	}
	owner, _ := database.GetOwner(ownerId)

	tenantId, err := strconv.Atoi(readValue(reader, "Tenant id"))
	if err != nil {
		panic(err)
	}
	tenant, _ := database.GetTenant(tenantId)

	apartmentId, err := strconv.Atoi(readValue(reader, "Apartment id"))
	if err != nil {
		panic(err)
	}
	apartment, _ := database.GetApartment(apartmentId)

	return domain.LeaseContract{
		From:      from,
		To:        to,
		Owner:     owner,
		Tenant:    tenant,
		Apartment: apartment,
	}
}

func readValue(reader *bufio.Reader, message string) string {
	fmt.Printf("%s > ", message)
	command, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(command, "\n")
}
