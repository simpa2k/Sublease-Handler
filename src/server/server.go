package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"subLease/src/server/database"
)

type Server struct {
	database database.Database
}

func Create(database database.Database) Server {
	return Server{
		database: database,
	}
}

func (s Server) Run() {
	r := s.SetupRouter()
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func (s Server) SetupRouter() *mux.Router {
	r := mux.NewRouter()

	s.setupApartmentRoutes(r)
	s.setupOwnerRoutes(r)
	s.setupTenantRoutes(r)
	s.setupLeaseContractRoutes(r)

	return r
}

func (s Server) setupLeaseContractRoutes(r *mux.Router) {
	r.HandleFunc("/lease_contract", getLeaseContractsHandler(s.database)).Methods("GET")
	r.HandleFunc("/lease_contract/{id}", getLeaseContractHandler(s.database)).Methods("GET")
	r.HandleFunc("/lease_contract", createLeaseContractHandler(s.database)).Methods("POST")
	r.HandleFunc("/lease_contract", updateLeaseContractHandler(s.database)).Methods("PUT")
	r.HandleFunc("/lease_contract/{id}", deleteLeaseContractHandler(s.database)).Methods("DELETE")
}

func (s Server) setupTenantRoutes(r *mux.Router) {
	r.HandleFunc("/tenant", getTenantsHandler(s.database)).Methods("GET")
	r.HandleFunc("/tenant/{id}", getTenantHandler(s.database)).Methods("GET")
	r.HandleFunc("/tenant", createTenantHandler(s.database)).Methods("POST")
	r.HandleFunc("/tenant/{id}", updateTenantHandler(s.database)).Methods("PUT")
	r.HandleFunc("/tenant/{id}", deleteTenantHandler(s.database)).Methods("DELETE")
}

func (s Server) setupOwnerRoutes(r *mux.Router) {
	r.HandleFunc("/owner", getOwnersHandler(s.database)).Methods("GET")
	r.HandleFunc("/owner/{id}", getOwnerHandler(s.database)).Methods("GET")
	r.HandleFunc("/owner", createOwnerHandler(s.database)).Methods("POST")
	r.HandleFunc("/owner", updateOwnerHandler(s.database)).Methods("PUT")
	r.HandleFunc("/owner/{id}", deleteOwnerHandler(s.database)).Methods("DELETE")
}

func (s Server) setupApartmentRoutes(r *mux.Router) {
	r.HandleFunc("/apartment", getApartmentsHandler(s.database)).Methods("GET")
	r.HandleFunc("/apartment/{id}", getApartmentHandler(s.database)).Methods("GET")
	r.HandleFunc("/apartment", createApartmentHandler(s.database)).Methods("POST")
	r.HandleFunc("/apartment/{id}", updateApartmentHandler(s.database)).Methods("PUT")
	r.HandleFunc("/apartment/{id}", deleteApartmentHandler(s.database)).Methods("DELETE")
}
