package codeGeneration

import "fmt"

type Modifier int

const (
	Unmodified Modifier = iota
	Slice      Modifier = iota
)

type Field struct {
	Identifier          string
	TypeDeclaration     string
	IsDomainEntity      bool
	IsComplexType       bool
	TypeModifier        Modifier
	FromStringConverter func(string) string
	HandlerFunction     func(string, string) string
}

func CreateIdField(identifier string) Field {
	return Field{
		Identifier:      identifier,
		TypeDeclaration: "int",
		IsDomainEntity:  false,
		IsComplexType:   false,
		TypeModifier:    Unmodified,
		FromStringConverter: func(stringValueIdentifier string) string {
			return fmt.Sprintf("return strconv.Atoi(%s)", stringValueIdentifier)
		},
		HandlerFunction: func(receiverIdentifier string, parsedValueIdentifier string) string {
			return fmt.Sprintf("%s = %s", receiverIdentifier, parsedValueIdentifier)
		},
	}
}

func CreateIntField(identifier string) Field {
	return Field{
		Identifier:      identifier,
		TypeDeclaration: "int",
		IsDomainEntity:  false,
		IsComplexType:   false,
		TypeModifier:    Unmodified,
		FromStringConverter: func(stringValueIdentifier string) string {
			return fmt.Sprintf("return strconv.Atoi(%s)", stringValueIdentifier)
		},
		HandlerFunction: assignmentToUpdateStructHandlerFunction(identifier),
	}
}

func CreateComplexField(identifier string, typeDeclaration string, isDomainEntity bool, typeModifier Modifier) Field {
	decapitalizedIdentifier := Decapitalize(identifier)
	return Field{
		Identifier:      identifier,
		TypeDeclaration: typeDeclaration,
		IsDomainEntity:  isDomainEntity,
		IsComplexType:   true,
		TypeModifier:    typeModifier,
		FromStringConverter: func(stringValueIdentifier string) string {
			return fmt.Sprintf("var %s %s; err := json.NewDecoder(strings.NewReader(%s)).Decode(&%s); return %s, err", decapitalizedIdentifier, typeDeclaration, stringValueIdentifier, decapitalizedIdentifier, decapitalizedIdentifier)
		},
		HandlerFunction: assignmentToUpdateStructHandlerFunction(identifier),
	}
}

func CreateTimeField(identifier string) Field {
	return Field{
		Identifier:      identifier,
		TypeDeclaration: "time.Time",
		IsDomainEntity:  false,
		IsComplexType:   false,
		TypeModifier:    Unmodified,
		FromStringConverter: func(stringValueIdentifier string) string {
			return fmt.Sprintf("return time.Parse(\"2006-01-02 15:04:05.999999999 -0700 MST\", %s)", stringValueIdentifier)
		},
		HandlerFunction: assignmentToUpdateStructHandlerFunction(identifier),
	}
}

func CreateDomainEntityField(typeDeclaration string) Field {
	return Field{
		Identifier:      typeDeclaration,
		TypeDeclaration: typeDeclaration,
		IsDomainEntity:  true,
		IsComplexType:   true,
		TypeModifier:    Unmodified,
		FromStringConverter: func(stringValueIdentifier string) string {
			return fmt.Sprintf("return strconv.Atoi(%s)", stringValueIdentifier)
		}, // TODO: This will work since domain entities will always be passed as an id in query parameters, but it looks a bit strange here.
		HandlerFunction: assignmentToUpdateStructHandlerFunction(typeDeclaration),
	}
}

func CreateDomainEntitySliceField(domainEntity string) Field {
	decapitalizedIdentifier := Decapitalize(domainEntity)
	pluralIdentifier := domainEntity + "s"
	return Field{
		Identifier:      pluralIdentifier,
		TypeDeclaration: "[]" + domainEntity,
		IsDomainEntity:  true,
		IsComplexType:   true,
		TypeModifier:    Slice,
		FromStringConverter: func(stringValueIdentifier string) string {
			return fmt.Sprintf("var %s []int; err := json.NewDecoder(strings.NewReader(%s)).Decode(&%s); return %s, err", decapitalizedIdentifier, stringValueIdentifier, decapitalizedIdentifier, decapitalizedIdentifier)
		},
		HandlerFunction: assignmentToUpdateStructHandlerFunction(pluralIdentifier),
	}
}

func CreateStringField(identifier string) Field {
	return Field{
		Identifier:          identifier,
		TypeDeclaration:     "string",
		IsDomainEntity:      false,
		IsComplexType:       false,
		TypeModifier:        Unmodified,
		FromStringConverter: func(stringValueIdentifier string) string { return fmt.Sprintf("return %s, nil", stringValueIdentifier) }, // TODO: This isn't very readable. It will be formatted as 'return s, nil' which will work, but not apparent that it will be called as a function.
		HandlerFunction:     assignmentToUpdateStructHandlerFunction(identifier),
	}
}

func assignmentToUpdateStructHandlerFunction(identifier string) func(string, string) string {
	return func(receiverIdentifier string, parsedValueIdentifier string) string {
		return fmt.Sprintf("%s.%s = &%s", receiverIdentifier, Capitalize(identifier), parsedValueIdentifier)
	}
}

