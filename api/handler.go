package api

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var taxType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tax",
	Fields: graphql.Fields{
		"amount":  &graphql.Field{Type: graphql.Float},
		"taxRate": &graphql.Field{Type: graphql.Float},
		"type":    &graphql.Field{Type: graphql.String},
	},
})

var salaryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Salary",
	Fields: graphql.Fields{
		"grossSalary":      &graphql.Field{Type: graphql.Float},
		"netSalary":        &graphql.Field{Type: graphql.Float},
		"anualGrossSalary": &graphql.Field{Type: graphql.Float},
		"anualNetSalary":   &graphql.Field{Type: graphql.Float},
		"taxes":            &graphql.Field{Type: graphql.NewList(taxType)},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"calculateTax": &graphql.Field{
			Type: salaryType,
			Args: graphql.FieldConfigArgument{
				"country":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"salaryPerMonth": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				country := p.Args["country"].(string)
				salary := p.Args["salaryPerMonth"].(float64)
				resolver := &Resolver{}
				return resolver.CalculateTax(country, salary)
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func NewGraphQLHandler() *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // Enables GraphiQL UI
	})
}
