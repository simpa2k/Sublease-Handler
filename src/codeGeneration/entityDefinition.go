package codeGeneration

type EntityDefinition struct {
	Entity       string
	Abbreviation string
	Fields       []Field
}

func GetDomainEntities() []EntityDefinition {
	return []EntityDefinition{
		{"apartment", "a", []Field{
			CreateIdField("Id"),
			CreateIntField("Number"),
			CreateComplexField("Address", "address.Address", false, Unmodified),
		}},
		{"lease_contract", "lc", []Field{
			CreateIdField("Id"),
			CreateTimeField("From"),
			CreateTimeField("To"),
			CreateDomainEntityField("Owner"),
			CreateDomainEntityField("Tenant"),
			CreateDomainEntityField("Apartment"),
		}},
		{"owner", "o", []Field{
			CreateIdField("Id"),
			CreateStringField("FirstName"),
			CreateStringField("LastName"),
			CreateComplexField("SocialSecurityNumber", "socialSecurityNumber.SocialSecurityNumber", false, Unmodified),
			CreateDomainEntitySliceField("Apartment"),
		}},
		{"tenant", "t", []Field{
			CreateIdField("Id"),
			CreateStringField("FirstName"),
			CreateStringField("LastName"),
			CreateComplexField("SocialSecurityNumber", "socialSecurityNumber.SocialSecurityNumber", false, Unmodified),
		}},
	}
}
