package main

import (
	"os"
	"strings"
	"io/ioutil"
	"text/template"
	"fmt"
)

type Operation struct {
	Entity string
}

type EntityDefinition struct {
	Entity string
	Abbreviation string
	Fields map[string]string
}

func main() {
	domainEntities := map[string]map[string]string {
		"apartment": {
			"Id": "int",
			"Number": "int",
			"Address": "address.Address",
		},
		"lease_contract": {
			"Id": "int",
			"From": "time.Time",
			"To": "time.Time",
			"Owner": "Owner",
			"Tenant": "Tenant",
			"Apartment": "Apartment",
		},
		"owner": {
			"Id": "int",
			"FirstName": "string",
			"LastName": "string",
			"SocialSecurityNumber": "socialSecurityNumber.SocialSecurityNumber",
			"Apartments": "[]Apartment",
		},
		"tenant": {
			"Id": "int",
			"FirstName": "string",
			"LastName": "string",
			"SocialSecurityNumber": "socialSecurityNumber.SocialSecurityNumber",
		},
	}

	funcMap := template.FuncMap {
		"Capitalize": func(s string) string { return applyToFirstCharacter(s, strings.ToUpper) },
		"Decapitalize": func(s string) string { return applyToFirstCharacter(s, strings.ToLower) },
		"CommaSeparate": commaSeparate,
		"BuildEquals": buildEquals,
		"IsDomainEntity": isDomainEntity,
	}

	generateMockOperations(domainEntities, funcMap, os.Args[1], os.Args[2])
	generateEntities(domainEntities, funcMap, os.Args[1], os.Args[3])
}

func generateMockOperations(domainEntities map[string]map[string]string, funcMap template.FuncMap, templateRoot string, outputPath string) {
	for entity := range domainEntities {
		op := Operation{toPascalCase(entity)}
		templateText, err := ioutil.ReadFile(templateRoot + "operations.tmpl")
		check(err)

		t, err := template.New("operation").Funcs(funcMap).Parse(string(templateText))
		check(err)

		f, err := os.Create(fmt.Sprintf("%s%sOperations.go", outputPath, toCamelCase(entity)))
		check(err)
		err = t.Execute(f, op)
		check(err)
	}
}

func generateEntities(domainEntities map[string]map[string]string, funcMap template.FuncMap, templateRoot string, outputPath string) {

	for entityName, fields := range domainEntities {
		entityDefinition := EntityDefinition {
			toPascalCase(entityName),
			abbreviate(entityName),
			fields,
		}

		t, err := template.New("domain_entity.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "domain_entity.tmpl")
		check(err)
		err = t.Execute(os.Stdout, entityDefinition)
		check(err)

		f, err := os.Create(fmt.Sprintf("%s%s.go", outputPath, toCamelCase(entityName)))
		check(err)
		err = t.Execute(f, entityDefinition)
		check(err)

		t, err = template.New("entity_update.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "entity_update.tmpl")
		check(err)

		f, err = os.Create(fmt.Sprintf("%s%sUpdate.go", outputPath, toCamelCase(entityName)))
		check(err)
		err = t.Execute(f, entityDefinition)
		check(err)
	}
}

func toPascalCase(snakeCased string) string {
	spaced := strings.Replace(snakeCased, "_", " ", -1)
	upperCased := strings.Title(spaced)
	return strings.Replace(upperCased, " ", "", -1)
}

func toCamelCase(snakeCased string) string {
	words := strings.Split(snakeCased, "_")
	camelCased := string(words[0])
	for _, word := range words[1:] {
		camelCased += strings.Title(word)
	}

	return camelCased
}

func abbreviate(snakeCased string) string {
	var abbreviated string
	words := strings.Split(snakeCased, "_")
	for _, word := range words {
		abbreviated += string(word[0])
	}

	return abbreviated
}

func commaSeparate(values map[string]string) string {
	var separated string
	first := true
	for key, value := range values {
		if first {
			first = false
		} else {
			separated += ", "
		}
		separated += fmt.Sprintf("%s %s", key, value)
	}
	return separated
}

func buildEquals(abbreviation string, values map[string]string) string {
	var separated string
	first := true
	for key := range values {
		if first {
			first = false
		} else {
			separated += " && "
		}
		separated += fmt.Sprintf("%s.%s == other.%s", abbreviation, key, key)
	}
	return separated
}

func isDomainEntity (typeDeclaration string) bool {
	domainEntities := map[string]struct{} {
		"Apartment": {},
		"LeaseContract": {},
		"Owner": {},
		"Tenant": {},
	}

	_, found := domainEntities[typeDeclaration]
	return found
}

func applyToFirstCharacter(s string, operation func(string) string) string {
	return operation(string(s[0])) + s[1:]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}