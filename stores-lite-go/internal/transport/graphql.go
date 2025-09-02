package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"
	"stores-lite/internal/domain"
	"stores-lite/internal/service"
)

func RegisterGraphQL(r *chi.Mux, svc *service.Service) {
	productType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"name":       &graphql.Field{Type: graphql.String},
			"priceCents": &graphql.Field{Type: graphql.Int},
			"stock":      &graphql.Field{Type: graphql.Int},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"products": &graphql.Field{
				Type: graphql.NewList(productType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return svc.ListProducts(p.Context)
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})

	r.Get("/graphiql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><form method="post" action="/graphql">
<textarea name="query" rows="10" cols="80">query { products { id name priceCents stock } }</textarea><br/>
<button type="submit">Run</button></form></body></html>`))
	})

	r.Post("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var body struct{ Query string `json:"query"` }
		_ = json.NewDecoder(r.Body).Decode(&body)
		result := graphql.Do(graphql.Params{Schema: schema, RequestString: body.Query, Context: r.Context()})
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})
}
