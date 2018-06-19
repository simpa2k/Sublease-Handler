package main

import (
	"fmt"
	"os"
	"subLease/src/codeGeneration"
	"subLease/src/codeGeneration/server/queryValueRetrieval"
	"text/template"
)

func main() {
	domainEntities := codeGeneration.GetDomainEntities()

	funcMap := template.FuncMap{
		"Capitalize":                 codeGeneration.Capitalize,
		"Decapitalize":               codeGeneration.Decapitalize,
		"DecapitalizeSlice":          codeGeneration.DecapitalizeSlice,
		"PascalCase":                 codeGeneration.ToPascalCase,
		"CamelCase":                  codeGeneration.ToCamelCase,
		"CamelCaseToNormalText":      codeGeneration.CamelCaseToNormalText,
		"CommaSeparate":              codeGeneration.CommaSeparate,
		"BuildEquals":                codeGeneration.BuildEquals,
		"IsSliceType":                codeGeneration.IsSliceType,
		"RequiresQuery":              codeGeneration.RequiresQuery,
		"NoIdField":                  codeGeneration.NoIdField,
		"RetrieveFromQueryParameter": codeGeneration.RetrieveFromQueryParameter,
		"FormatTypeDeclaration":      queryValueRetrieval.FormatTypeDeclaration,
	}

	generate(domainEntities, funcMap)
}

func generate(domainEntities []codeGeneration.EntityDefinition, funcMap template.FuncMap) {
	queryValueRetrievalGenerator := queryValueRetrieval.Create()
	for _, entityDefinition := range domainEntities {
		generateDatabaseOperations(entityDefinition, funcMap, os.Args[1], os.Args[4])
		generateInMemoryDatabaseOperations(entityDefinition, funcMap, os.Args[1], os.Args[2])
		generateDomainEntity(entityDefinition, funcMap, os.Args[1], os.Args[3])
		generateDomainEntityUpdate(entityDefinition, funcMap, os.Args[1], os.Args[4])
		generateHandlers(entityDefinition, &queryValueRetrievalGenerator, funcMap, os.Args[1], os.Args[5])
		generateCommands(entityDefinition, funcMap, os.Args[1], os.Args[6])
	}
	queryValueRetrievalGenerator.Generate(os.Args[1], os.Args[5])
	generateDatabase(domainEntities, funcMap, os.Args[1], os.Args[4])
}

func generateDatabase(domainEntities []codeGeneration.EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	t, err := template.New("database.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "server/database.tmpl")
	check(err)
	f, err := os.Create(fmt.Sprintf("%sdatabase.go", outputPath))
	check(err)
	err = t.Execute(f, domainEntities)
	check(err)
}

func generateDatabaseOperations(domainEntity codeGeneration.EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	generateOperations(domainEntity, funcMap, templateRoot+"/server/", "database_operations.tmpl", outputPath)
}

func generateInMemoryDatabaseOperations(domainEntity codeGeneration.EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	generateOperations(domainEntity, funcMap, templateRoot+"/inMemoryDatabase/", "in_memory_database_operations.tmpl", outputPath)
}

func generateOperations(domainEntity codeGeneration.EntityDefinition, funcMap template.FuncMap, templateRoot string, templateName string, outputPath string) {
	t, err := template.New(templateName).Funcs(funcMap).ParseFiles(templateRoot + templateName)
	check(err)

	f, err := os.Create(fmt.Sprintf("%s%sOperations.go", outputPath, codeGeneration.ToCamelCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func generateDomainEntity(domainEntity codeGeneration.EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	t, err := template.New("domain_entity.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "server/domain_entity.tmpl")
	check(err)

	f, err := os.Create(fmt.Sprintf("%s%s.go", outputPath, codeGeneration.ToCamelCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func generateDomainEntityUpdate(domainEntity codeGeneration.EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	replaceDomainEntityTypesWithInt(domainEntity)

	t, err := template.New("entity_update.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "server/entity_update.tmpl")
	check(err)

	f, err := os.Create(fmt.Sprintf("%s%sUpdate.go", outputPath, codeGeneration.ToCamelCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func generateHandlers(domainEntity codeGeneration.EntityDefinition, queryValueRetrievalGenerator *queryValueRetrieval.QueryValueRetrievalGenerator, funcMap template.FuncMap, templateRoot string, outputPath string) {
	for _, field := range domainEntity.Fields {
		queryValueRetrievalGenerator.AddType(field.TypeDeclaration)
	}

	replaceDomainEntityTypesWithInt(domainEntity)
	t, err := template.New("handlers.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "server/handlers.tmpl")
	check(err)

	f, err := os.Create(fmt.Sprintf("%s%sHandlers.go", outputPath, codeGeneration.ToCamelCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func generateCommands(domainEntity codeGeneration.EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	generatePostCommand(domainEntity, funcMap, templateRoot, outputPath)
	generateListCommand(domainEntity, funcMap, templateRoot, outputPath)
}

func generatePostCommand(domainEntity codeGeneration.EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	t, err := template.New("postCommand.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "client/postCommand.tmpl")
	check(err)

	f, err := os.Create(fmt.Sprintf("%spost%sCommand.go", outputPath, codeGeneration.ToPascalCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func generateListCommand(domainEntity codeGeneration.EntityDefinition, funcMap template.FuncMap, templateRoot string, outputPath string) {
	t, err := template.New("listCommand.tmpl").Funcs(funcMap).ParseFiles(templateRoot + "client/listCommand.tmpl")
	check(err)

	f, err := os.Create(fmt.Sprintf("%slist%ssCommand.go", outputPath, codeGeneration.ToPascalCase(domainEntity.Entity)))
	check(err)
	err = t.Execute(f, domainEntity)
	check(err)
}

func replaceDomainEntityTypesWithInt(domainEntity codeGeneration.EntityDefinition) {
	for i, field := range domainEntity.Fields {
		if field.IsDomainEntity { // Updates only use ids for referencing domain entities
			if field.TypeModifier == codeGeneration.Slice {
				field.TypeDeclaration = "[]int"
			} else {
				field.TypeDeclaration = "int"
			}
		}
		domainEntity.Fields[i] = field
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
