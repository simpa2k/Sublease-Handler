package queryValueRetrieval

import (
	"os"
	"fmt"
	"text/template"
	"strings"
)

type QueryValueRetrievalGenerator struct {
	typesSet map[string]struct{}
	typesSlice []string
	funcMap template.FuncMap
}

func Create() (QueryValueRetrievalGenerator) {
	return QueryValueRetrievalGenerator{
		typesSet: make(map[string]struct{}),
		funcMap: template.FuncMap{
			"FormatTypeDeclaration": formatTypeDeclaration,
		},
	}
}

func formatTypeDeclaration(typeDeclaration string) string {
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

func (qvrg *QueryValueRetrievalGenerator) AddType(t string) {
	if _, found := qvrg.typesSet[t]; !found {
		qvrg.typesSet[t] = struct{}{}
		qvrg.typesSlice = append(qvrg.typesSlice, t)
	}
}

func (qvrg QueryValueRetrievalGenerator) Generate(templateRoot string, outputPath string) {
	t, err := template.New("query_value_retrieval.tmpl").Funcs(qvrg.funcMap).ParseFiles(templateRoot + "queryValueRetrieval/query_value_retrieval.tmpl")
	check(err)

	f, err := os.Create(fmt.Sprintf("%s%s.go", outputPath, "queryValueRetrieval"))
	check(err)
	err = t.Execute(f, qvrg.typesSlice)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
