package codeGeneration

import (
	"fmt"
	"regexp"
	"strings"
)

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

func CamelCaseToNormalText(camelCased string) string {
	regex := regexp.MustCompile("([A-Z])")
	return strings.ToLower(regex.ReplaceAllString(camelCased, " $1"))
}

func CommaSeparate(values []Field) string {
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

func BuildEquals(entityDefinition EntityDefinition) string {
	var separated string
	first := true
	for _, field := range entityDefinition.Fields {
		if field.TypeModifier != Slice { // TODO: Handle this
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

func IsSliceType(field Field) bool {
	return field.TypeModifier == Slice
}

func RequiresQuery(field Field) bool {
	return field.IsDomainEntity && field.TypeModifier != Slice
}

func NoIdField(fields []Field) []Field {
	var filtered []Field
	for _, field := range fields {
		if field.Identifier != "Id" {
			filtered = append(filtered, field)
		}
	}
	return filtered
}

func DecapitalizeSlice(s []Field) []Field {
	var decapitalized []Field
	for _, field := range s {
		field.Identifier = ApplyToFirstCharacter(field.Identifier, strings.ToLower)
		decapitalized = append(decapitalized, field)
	}

	return decapitalized
}

func Capitalize(s string) string {
	return ApplyToFirstCharacter(s, strings.ToUpper)
}

func Decapitalize(s string) string {
	return ApplyToFirstCharacter(s, strings.ToLower)
}

func ApplyToFirstCharacter(s string, operation func(string) string) string {
	return operation(string(s[0])) + s[1:]
}

func RetrieveFromQueryParameter(valuesIdentifier string, dateLayout string, field Field) string {
	retrieval := ""
	fieldName := ApplyToFirstCharacter(field.Identifier, strings.ToLower)

	if field.TypeDeclaration == "string" {
		retrieval = fmt.Sprintf("%s := %s.GetDomainEntities(\"%s\")", fieldName, valuesIdentifier, fieldName)
	} else if field.TypeDeclaration == "int" {
		retrieval = fmt.Sprintf("%s, _ := strconv.Atoi(%s.GetDomainEntities(\"%s\"))", fieldName, valuesIdentifier, fieldName)
	} else if field.TypeDeclaration == "time.Time" {
		retrieval = fmt.Sprintf("%s, err := time.Parse(%s, %s.GetDomainEntities(\"%s\")); if err != nil { panic(err) }", fieldName, dateLayout, valuesIdentifier, fieldName)
	} else if field.TypeModifier == Slice || field.IsComplexType {
		retrieval = fmt.Sprintf("var %s %s; _ = json.NewDecoder(strings.NewReader(%s.GetDomainEntities(\"%s\"))).Decode(&%s)", fieldName, field.TypeDeclaration, valuesIdentifier, fieldName, fieldName)
	}

	return retrieval
}
