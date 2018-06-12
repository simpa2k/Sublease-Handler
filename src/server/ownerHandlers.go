// Generated by text/template; DO NOT EDIT
package server




func getOwnersHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(database.GetOwners())
	}
}

func getOwnerHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		if owner, found := database.GetOwner(id); found {
			json.NewEncoder(w).Encode(owner)
		}
	}
}

func createOwnerHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var owner domain.Owner
		_ = json.NewDecoder(r.Body).Decode(&owner)

		json.NewEncoder(w).Encode(database.CreateOwner(owner))
	}
}

func updateOwnerHandler(db database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	    queryValues := r.URL.Query()
        strconv.Atoi(id)
        s, nil
        s, nil
        var socialSecurityNumber socialSecurityNumber.SocialSecurityNumber; _ = json.NewDecoder(strings.NewReader(socialSecurityNumber))).Decode(&socialSecurityNumber)
        var apartments []Apartment; _ = json.NewDecoder(strings.NewReader(apartments))).Decode(&apartments)

		ownerUpdate := database.OwnerUpdate{
            FirstName: &firstName,
            LastName: &lastName,
            SocialSecurityNumber: &socialSecurityNumber,
            Apartments: &apartments,
		}

		updatedOwner, foundOwnerWithId := db.UpdateOwner(id, ownerUpdate)
		if foundOwnerWithId {
			json.NewEncoder(w).Encode(updatedOwner)
		} else {
            http.Error(w, "No owner with that id was found.", http.StatusBadRequest)
		}
	}
}

func deleteOwnerHandler(database database.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		if owner, found := database.DeleteOwner(id); found {
			json.NewEncoder(w).Encode(owner)
		}
	}
}
