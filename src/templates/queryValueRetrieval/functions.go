package queryValueRetrieval

import "strings"

func FormatTypeDeclaration(typeDeclaration string) string {
	formatted := removePackageQualifier(typeDeclaration)
	formatted = handleSliceSymbols(formatted)

	return strings.Title(formatted)
}

func removePackageQualifier(typeDeclaration string) string {
	parts := strings.Split(typeDeclaration, ".")
	var typePart string
	if len(parts) > 1 {
		typePart = parts[1]
	} else {
		typePart = parts[0]
	}
	return typePart
}

func handleSliceSymbols(typeDeclaration string) string {
	if len(typeDeclaration) > 3 && typeDeclaration[0:2] == "[]" {
		typeDeclaration = typeDeclaration[2:] + "Slice"
	}

	return typeDeclaration
}
