package queryValueRetrieval

import (
	"fmt"
	"os"
	"text/template"
)

type QueryValueRetrievalGenerator struct {
	typesSet   map[string]struct{}
	typesSlice []string
	funcMap    template.FuncMap
}

func Create() QueryValueRetrievalGenerator {
	return QueryValueRetrievalGenerator{
		typesSet: make(map[string]struct{}),
		funcMap: template.FuncMap{
			"FormatTypeDeclaration": FormatTypeDeclaration,
		},
	}
}

func (qvrg *QueryValueRetrievalGenerator) AddType(t string) {
	if _, found := qvrg.typesSet[t]; !found {
		qvrg.typesSet[t] = struct{}{}
		qvrg.typesSlice = append(qvrg.typesSlice, t)
	}
}

func (qvrg QueryValueRetrievalGenerator) Generate(templateRoot string, outputPath string) {
	t, err := template.New("query_value_retrieval.tmpl").Funcs(qvrg.funcMap).ParseFiles(templateRoot + "server/queryValueRetrieval/query_value_retrieval.tmpl")
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
