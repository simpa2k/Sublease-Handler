package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"regexp"
	"subLease/src/templates/queryValueRetrieval"
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
	FromStringConverter string
}

func CreateIntField(identifier string) Field {
	return Field{
		Identifier: identifier,
		TypeDeclaration: "int",
		IsDomainEntity: false,
		IsComplexType: false,
		TypeModifier: unmodified,
		FromStringConverter: fmt.Sprintf("strconv.Atoi(%s)", decapitalize(identifier)),
	}
}

func CreateComplexField(identifier string, typeDeclaration string, isDomainEntity bool, typeModifier modifier) Field {
	decapitalizedIdentifier := decapitalize(identifier)
	return Field{
		Identifier: identifier,
		TypeDeclaration: typeDeclaration,
		IsDomainEntity: isDomainEntity,
		IsComplexType: true,
		TypeModifier: typeModifier,
		FromStringConverter: fmt.Sprintf("var %s %s; _ = json.NewDecoder(strings.NewReader(%s))).Decode(&%s)", decapitalizedIdentifier,  typeDeclaration, decapitalizedIdentifier, decapitalizedIdentifier),
	}
}

func CreateTimeField(identifier string) Field {
	return Field{
		Identifier: identifier,
		TypeDeclaration: "time.Time",
		IsDomainEntity: false,
		IsComplexType: false,
		TypeModifier: unmodified,
		FromStringConverter: fmt.Sprintf("time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", %s)", decapitalize(identifier)),
	}
}

func CreateDomainEntityField(typeDeclaration string) Field {
	return Field{
		Identifier: typeDeclaration,
		TypeDeclaration: typeDeclaration,
		IsDomainEntity: true,
		IsComplexType: true,
		TypeModifier: unmodified,
		FromStringConverter: fmt.Sprintf("strconv.Atoi(%s)", decapitalize(typeDeclaration)), // TODO: This will work since domain entities will always be passed as an id in query parameters, but it looks a bit strange here.
	}
}

func CreateStringField(identifier string) Field {
	return Field{
		Identifier: identifier,
		TypeDeclaration: "string",
		IsDomainEntity: false,
		IsComplexType: false,
		TypeModifier: unmodified,
		FromStringConverter: "s, nil", // TODO: This isn't very readable. It will be formatted as 'return s, nil' which will work, but not apparent that it will be called as a function.
	}
}

func main() {
	domainEntities := []EntityDefinition{
		{"apartment", "a", []Field{
			CreateIntField("Id"),
			CreateIntField("Number"),
			CreateComplexField("Address", "address.Address", false, unmodified),
		}},
		{"lease_contract", "lc", []Field{
			CreateIntField("Id"),
			CreateTimeField("From"),
			CreateTimeField("To"),
			CreateDomainEntityField("Owner"),
			CreateDomainEntityField("Tenant"),
			CreateDomainEntityField("Apartment"),
		}},
		{"owner", "o", []Field{
			CreateIntField("Id"),
			CreateStringField("FirstName"),
			CreateStringField("LastName"),
			CreateComplexField("SocialSecurityNumber", "socialSecurityNumber.SocialSecurityNumber", false, unmodified),
			CreateComplexField("Apartments", "[]Apartment", true, slice),
		}},
		{"tenant", "t", []Field{
			CreateIntField("Id"),
			CreateStringField("FirstName"),
			CreateStringField("LastName"),
			CreateComplexField("SocialSecurityNumber", "socialSecurityNumber.SocialSecurityNumber", false, unmodified),
		}},
	}

	funcMap := template.FuncMap{
		"Capitalize":                 func(s string) string { return applyToFirstCharacter(s, strings.ToUpper) },
		"Decapitalize":               decapitalize,
		"DecapitalizeSlice":          decapitalizeSlice,
		"PascalCase":                 ToPascalCase,
		"CamelCase":                  ToCamelCase,
		"CamelCaseToNormalText":	  camelCaseToNormalText,
		"CommaSeparate":              commaSeparate,
		"BuildEquals":                buildEquals,
		"IsSliceType":                isSliceType,
		"RequiresQuery":              requiresQuery,
		"NoIdField":               noIdField,
		"RetrieveFromQueryParameter": retrieveFromQueryParameter,
	}

	generate(domainEntities, funcMap)
}

func generate(domainEntities []EntityDefinition, funcMap template.FuncMap) {
	queryValueRetrievalGenerator := queryValueRetrieval.Create()
	for _, entityDefinition := range domainEntities {
		generateDatabaseOperations(entityDefinition, funcMap, os.Args[1], os.Args[4])
		generateMockDatabaseOperations(entityDefinition, funcMap, os.Args[1], os.Args[2])
		generateDomainEntity(entityDefinition, funcMap, os.Args[1], os.Args[3])
		generateDomainEntityUpdate(entityDefinition, funcMap, os.Args[1], os.Args[4])
		generateHandlers(entityDefinition, &queryValueRetrievalGenerator, funcMap, os.Args[1], os.Args[5])
	}
	queryValueRetrievalGenerator.Generate(os.Args[1], os.Args[5])
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
	replaceDomainEntityTypesWithInt(domainEntity)

	t, err := template.New("entity_update.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "entity_update.tmpl")
	check(err)

	f, err := os.Create(fmt.Sprintf("%s%sUpdate.go", outputPath, ToCamelCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func generateHandlers(domainEntity EntityDefinition, queryValueRetrievalGenerator *queryValueRetrieval.QueryValueRetrievalGenerator, funcMap template.FuncMap, templateRoot string, outputPath string) {
	for _, field := range domainEntity.Fields {
		queryValueRetrievalGenerator.AddType(field.TypeDeclaration)
	}

	replaceDomainEntityTypesWithInt(domainEntity)
	t, err := template.New("handlers.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "handlers.tmpl")
	check(err)

	f, err := os.Create(fmt.Sprintf("%s%sHandlers.go", outputPath, ToCamelCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func replaceDomainEntityTypesWithInt(domainEntity EntityDefinition) {
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

func camelCaseToNormalText(camelCased string) string {
	regex := regexp.MustCompile("([A-Z])")
	return strings.ToLower(regex.ReplaceAllString(camelCased, " $1"))
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

func noIdField(fields []Field) []Field {
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

func decapitalize(s string) string {
	return applyToFirstCharacter(s, strings.ToLower)
}

func applyToFirstCharacter(s string, operation func(string) string) string {
	return operation(string(s[0])) + s[1:]
}

func retrieveFromQueryParameter(valuesIdentifier string, dateLayout string, field Field) string {
	/*retrieval := ""
	fieldName := applyToFirstCharacter(field.Identifier, strings.ToLower)

	if field.TypeDeclaration == "string" {
		retrieval = fmt.Sprintf("%s := %s.Get(\"%s\")", fieldName, valuesIdentifier, fieldName)
	} else if field.TypeDeclaration == "int" {
		retrieval = fmt.Sprintf("%s, _ := strconv.Atoi(%s.Get(\"%s\"))", fieldName, valuesIdentifier, fieldName)
	} else if field.TypeDeclaration == "time.Time" {
		retrieval = fmt.Sprintf("%s, err := time.Parse(%s, %s.Get(\"%s\")); if err != nil { panic(err) }", fieldName, dateLayout, valuesIdentifier, fieldName)
	} else if field.TypeModifier == slice || field.IsComplexType {
		retrieval = fmt.Sprintf("var %s %s; _ = json.NewDecoder(strings.NewReader(%s.Get(\"%s\"))).Decode(&%s)", fieldName, field.TypeDeclaration, valuesIdentifier, fieldName, fieldName)
	}

	return retrieval*/
	return field.FromStringConverter
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
