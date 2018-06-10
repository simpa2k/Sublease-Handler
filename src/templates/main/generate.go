package main

import (
	"os"
	"strings"
	"io/ioutil"
	"text/template"
	"fmt"
	"regexp"
)

type Operation struct {
	Entity string
}

type EntityDefinition struct {
	Entity string
	Abbreviation string
	//Fields map[string]string
	Fields []Field
}

type modifier int
const (
	unmodified modifier = iota
	slice	   modifier = iota
)

type Field struct {
	Identifier      string
	TypeDeclaration string
	IsDomainEntity  bool
	IsComplexType   bool
	TypeModifier    modifier
}

func main() {
	domainEntities := []EntityDefinition {
		{"apartment", "a", []Field {
			{"Id", "int", false, false, unmodified},
			{"Number","int", false, false, unmodified},
			{"Address", "address.Address",false, true, unmodified},
		}},
		{"lease_contract", "lc", []Field {
			{"Id", "int", false, false, unmodified},
			{"From", "time.Time", false, true, unmodified},
			{"To", "time.Time", false, true, unmodified},
			{"Owner", "Owner", true, true, unmodified},
			{"Tenant", "Tenant", true, true, unmodified},
			{"Apartment", "Apartment", true, true, unmodified},
		}},
		{"owner", "o", []Field {
			{"Id", "int", false, false, unmodified},
			{"FirstName", "string", false, false, unmodified},
			{"LastName", "string", false, false, unmodified},
			{"SocialSecurityNumber", "socialSecurityNumber.SocialSecurityNumber", false, true, unmodified},
			{"Apartments", "[]Apartment", true, true, slice},
		}},
		{"tenant", "t", []Field {
			{"Id","int", false, false, unmodified},
			{"FirstName","string", false, false, unmodified},
			{"LastName", "string", false, false, unmodified},
			{"SocialSecurityNumber", "socialSecurityNumber.SocialSecurityNumber", false, true, unmodified},
		}},
	}
	/*domainEntities := map[string]map[string]string {
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
	}*/

	funcMap := template.FuncMap{
		"Capitalize":                      func(s string) string { return applyToFirstCharacter(s, strings.ToUpper) },
		"Decapitalize":                    func(s string) string { return applyToFirstCharacter(s, strings.ToLower) },
		"DecapitalizeSlice":               decapitalizeSlice,
		"PascalCase":                      ToPascalCase,
		"CamelCase":                       ToCamelCase,
		"CommaSeparate":                   commaSeparate,
		"BuildEquals":                     buildEquals,
		"RemoveQualificationIfDomainType": removeQualificationIfDomainType,
		"IsSliceType":                     isSliceType,
		"RequiresQuery":				   requiresQuery,
		"FilterFieldsFromConstructor":	   filterFieldsFromConstructor,
	}

	generateMockOperations(domainEntities, funcMap, os.Args[1], os.Args[2])
	generateEntities(domainEntities, funcMap, os.Args[1], os.Args[3], os.Args[4])
}

func setComplexTypes(domainEntities map[string]map[string]string) map[string]struct{} {
	complexTypes := make(map[string]struct{})
	for key := range domainEntities {
		complexTypes[key] = struct{}{}
	}

	additionalComplexTypes := []string{"address", "social_security_number"}
	for _, additionalComplexType := range additionalComplexTypes {
		complexTypes[additionalComplexType] = struct{}{}
	}

	return complexTypes
}

func generateMockOperations(domainEntities []EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	for _, entity := range domainEntities {
		//op := Operation{ToPascalCase(entity.Entity)}
		op := Operation{entity.Entity}
		templateText, err := ioutil.ReadFile(templateRoot + "operations.tmpl")
		check(err)

		t, err := template.New("operation").Funcs(funcMap).Parse(string(templateText))
		check(err)

		f, err := os.Create(fmt.Sprintf("%s%sOperations.go", outputPath, ToCamelCase(entity.Entity)))
		check(err)
		err = t.Execute(f, op)
		check(err)
	}
}

func generateEntities(domainEntities []EntityDefinition, funcMap template.FuncMap, templateRoot string, entityOutputPath string, entityUpdateOutputPath string) {

	for _, entityDefinition := range domainEntities {

		t, err := template.New("domain_entity.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "domain_entity.tmpl")
		check(err)
		err = t.Execute(os.Stdout, entityDefinition)
		check(err)

		f, err := os.Create(fmt.Sprintf("%s%s.go", entityOutputPath, ToCamelCase(entityDefinition.Entity)))
		check(err)
		err = t.Execute(f, entityDefinition)
		check(err)

		for i, field := range entityDefinition.Fields {
			if field.IsDomainEntity {
				if field.TypeModifier == slice {
					field.TypeDeclaration = "[]int" // Updates only use ids for referencing domain entities
				} else {
					field.TypeDeclaration = "int"
				}
			}
			entityDefinition.Fields[i] = field
		}

		t, err = template.New("entity_update.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "entity_update.tmpl")
		check(err)

		f, err = os.Create(fmt.Sprintf("%s%sUpdate.go", entityUpdateOutputPath, ToCamelCase(entityDefinition.Entity)))
		check(err)
		err = t.Execute(f, entityDefinition)
		check(err)
	}
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

// From: https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6
func toSnakeCase(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func abbreviate(snakeCased string) string {
	var abbreviated string
	words := strings.Split(snakeCased, "_")
	for _, word := range words {
		abbreviated += string(word[0])
	}

	return abbreviated
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

func isComplexType(complexTypes map[string]struct{}, typeDeclaration string) bool {
	_, found := complexTypes[toSnakeCase(typeDeclaration)]
	return found
}

func isDomainEntity(domainEntities map[string]map[string]string, typeDeclaration string) bool {
	typeDeclaration = replaceNonAlphabetical(typeDeclaration)
	_, found := domainEntities[toSnakeCase(typeDeclaration)]
	return found
}

func replaceNonAlphabetical(s string) string {
	nonAlphabetical := regexp.MustCompile("[^A-Za-z]")
	return nonAlphabetical.ReplaceAllString(s, "")
}

func removeQualificationIfDomainType(s string) string {
	domainType := regexp.MustCompile("domain.")
	return domainType.ReplaceAllString(s, "")
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