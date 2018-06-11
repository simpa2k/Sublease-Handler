package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type EntityDefinition struct {
	Entity       string
	Abbreviation string
	Fields       []Field
}

type modifier int

const (
	unmodified modifier = iota
	slice      modifier = iota
)

type Field struct {
	Identifier      string
	TypeDeclaration string
	IsDomainEntity  bool
	IsComplexType   bool
	TypeModifier    modifier
}

func main() {
	domainEntities := []EntityDefinition{
		{"apartment", "a", []Field{
			{"Id", "int", false, false, unmodified},
			{"Number", "int", false, false, unmodified},
			{"Address", "address.Address", false, true, unmodified},
		}},
		{"lease_contract", "lc", []Field{
			{"Id", "int", false, false, unmodified},
			{"From", "time.Time", false, true, unmodified},
			{"To", "time.Time", false, true, unmodified},
			{"Owner", "Owner", true, true, unmodified},
			{"Tenant", "Tenant", true, true, unmodified},
			{"Apartment", "Apartment", true, true, unmodified},
		}},
		{"owner", "o", []Field{
			{"Id", "int", false, false, unmodified},
			{"FirstName", "string", false, false, unmodified},
			{"LastName", "string", false, false, unmodified},
			{"SocialSecurityNumber", "socialSecurityNumber.SocialSecurityNumber", false, true, unmodified},
			{"Apartments", "[]Apartment", true, true, slice},
		}},
		{"tenant", "t", []Field{
			{"Id", "int", false, false, unmodified},
			{"FirstName", "string", false, false, unmodified},
			{"LastName", "string", false, false, unmodified},
			{"SocialSecurityNumber", "socialSecurityNumber.SocialSecurityNumber", false, true, unmodified},
		}},
	}

	funcMap := template.FuncMap{
		"Capitalize":                  func(s string) string { return applyToFirstCharacter(s, strings.ToUpper) },
		"Decapitalize":                func(s string) string { return applyToFirstCharacter(s, strings.ToLower) },
		"DecapitalizeSlice":           decapitalizeSlice,
		"PascalCase":                  ToPascalCase,
		"CamelCase":                   ToCamelCase,
		"CommaSeparate":               commaSeparate,
		"BuildEquals":                 buildEquals,
		"IsSliceType":                 isSliceType,
		"RequiresQuery":               requiresQuery,
		"FilterFieldsFromConstructor": filterFieldsFromConstructor,
	}

	generate(domainEntities, funcMap)
}

func generate(domainEntities []EntityDefinition, funcMap template.FuncMap) {
	for _, entityDefinition := range domainEntities {
		generateDatabaseOperations(entityDefinition, funcMap, os.Args[1], os.Args[4])
		generateMockDatabaseOperations(entityDefinition, funcMap, os.Args[1], os.Args[2])
		generateDomainEntity(entityDefinition, funcMap, os.Args[1], os.Args[3])
		generateDomainEntityUpdate(entityDefinition, funcMap, os.Args[1], os.Args[4])
	}
	generateDatabase(domainEntities, funcMap, os.Args[1], os.Args[4])
}

func generateDatabase(domainEntities []EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	t, err := template.New("database.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "database.tmpl")
	check(err)
	f, err := os.Create(fmt.Sprintf("%sdatabase.go", outputPath))
	check(err)
	err = t.Execute(f, domainEntities)
	check(err)
}

func generateDatabaseOperations(domainEntity EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	generateOperations(domainEntity, funcMap, templateRoot, "database_operations.tmpl", outputPath)
}

func generateMockDatabaseOperations(domainEntity EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	generateOperations(domainEntity, funcMap, templateRoot, "mock_database_operations.tmpl", outputPath)
}

func generateOperations(domainEntity EntityDefinition, funcMap template.FuncMap, templateRoot string, templateName string, outputPath string) {
	t, err := template.New(templateName).Funcs(funcMap).ParseFiles(templateRoot + templateName)
	check(err)

	f, err := os.Create(fmt.Sprintf("%s%sOperations.go", outputPath, ToCamelCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func generateDomainEntity(domainEntity EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	t, err := template.New("domain_entity.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "domain_entity.tmpl")
	check(err)

	f, err := os.Create(fmt.Sprintf("%s%s.go", outputPath, ToCamelCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func generateDomainEntityUpdate(domainEntity EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	for i, field := range domainEntity.Fields {
		if field.IsDomainEntity { // Updates only use ids for referencing domain entities
			if field.TypeModifier == slice {
				field.TypeDeclaration = "[]int"
			} else {
				field.TypeDeclaration = "int"
			}
		}
		domainEntity.Fields[i] = field
	}

	t, err := template.New("entity_update.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "entity_update.tmpl")
	check(err)

	f, err := os.Create(fmt.Sprintf("%s%sUpdate.go", outputPath, ToCamelCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func ToPascalCase(snakeCased string) string {
	spaced := strings.Replace(snakeCased, "_", " ", -1)
	upperCased := strings.Title(spaced)
	return strings.Replace(upperCased, " ", "", -1)
}

func ToCamelCase(snakeCased string) string {
	words := strings.Split(snakeCased, "_")
	camelCased := strings.ToLower(string(words[0]))
	for _, word := range words[1:] {
		camelCased += strings.Title(word)
	}

	return camelCased
}

func commaSeparate(values []Field) string {
	var separated string
	first := true
	for _, field := range values {
		if first {
			first = false
		} else {
			separated += ", "
		}
		separated += fmt.Sprintf("%s %s", field.Identifier, field.TypeDeclaration)
	}
	return separated
}

func buildEquals(entityDefinition EntityDefinition) string {
	var separated string
	first := true
	for _, field := range entityDefinition.Fields {
		if field.TypeModifier != slice { // TODO: Handle this
			if first {
				first = false
			} else {
				separated += " && "
			}

			var stringToFormat string
			if field.IsDomainEntity {
				stringToFormat = "%s.%s.Equal(&other.%s)"
			} else if field.IsComplexType {
				stringToFormat = "%s.%s.Equal(other.%s)"
			} else {
				stringToFormat = "%s.%s == other.%s"
			}

			separated += fmt.Sprintf(stringToFormat, entityDefinition.Abbreviation, field.Identifier, field.Identifier)
		}
	}
	return separated
}

func isSliceType(field Field) bool {
	return field.TypeModifier == slice
}

func requiresQuery(field Field) bool {
	return field.IsDomainEntity && field.TypeModifier != slice
}

func filterFieldsFromConstructor(fields []Field) []Field {
	var filtered []Field
	for _, field := range fields {
		if field.Identifier != "Id" {
			filtered = append(filtered, field)
		}
	}
	return filtered
}

func decapitalizeSlice(s []Field) []Field {
	var decapitalized []Field
	for _, field := range s {
		field.Identifier = applyToFirstCharacter(field.Identifier, strings.ToLower)
		decapitalized = append(decapitalized, field)
	}

	return decapitalized
}

func applyToFirstCharacter(s string, operation func(string) string) string {
	return operation(string(s[0])) + s[1:]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
