package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Player holds player response
type Player struct {
	ID             int      `json:"int"`
	Name           string   `json:"name"`
	HighScore      int      `json:"highScore"`
	IsOnline       bool     `json:"isOnline"`
	Location       string   `json:"location"`
	LevelsUnlocked []string `json:"levelsUnlocked"`
}

var players = []Player{
	Player{ID: 123, Name: "Pablo", HighScore: 1100, IsOnline: true, Location: "Italy"},
	Player{ID: 230, Name: "Dora", HighScore: 2100, IsOnline: false, Location: "Germany"},
}

var playerObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"highScore": &graphql.Field{
				Type: graphql.String,
			},
			"isOnline": &graphql.Field{
				Type: graphql.Boolean,
			},
			"location": &graphql.Field{
				Type: graphql.String,
			},
			"levelsUnlocked": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

func main() {
	// Schema
	fields := graphql.Fields{
		"players": &graphql.Field{
			Type:        graphql.NewList(playerObject),
			Description: "All players",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return players, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
