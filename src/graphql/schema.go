package graphql

import (
	"fmt"
	"log"
	"net/http"

	gql "github.com/graph-gophers/graphql-go"
)

var SchemaString = `
schema {
	query: Query
	mutation: Mutation
}

type Query {}

type Mutation {
	createDownloadJob(playlist: String!): PlaylistJob!
}
` + PlaylistDef

// GetSchema Returns the parsed graphql schema that matches
// the concatenated schemaString found in the graphql package
func GetSchema() *gql.Schema {
	return gql.MustParseSchema(SchemaString, &Resolver{})
}

// Resolver is a golang-go resolver body for all of the graphql schema types,
// queries, and mutations any defined query or mutation must have a matching
// publicly exposed member method
type Resolver struct{}

// CreateDownloadJob will create a job for the playlist to be downloaded in the background
// a job will only be created after given url passes validation
func (r *Resolver) CreateDownloadJob(args *struct{ Playlist string }) (*PlaylistJobResolver, error) {
	// TODO: add validation
	// TODO: add database entry and task parsing
	log.Println("Creating download")

	resp, err := http.Get("https://www.youtube.com/playlist?list=" + args.Playlist)

	fmt.Printf("%+v \n", resp.Body)
	byteArr := make([]byte, 0)
	resp.Body.Read(byteArr)
	fmt.Printf("%+v \n", byteArr)

	if err != nil {
		return nil, err
	}
	job := &PlaylistJob{id: "1", url: args.Playlist}

	return &PlaylistJobResolver{job}, nil
}

// PlaylistJob will return the playlist job for a given ID
func (r *Resolver) PlaylistJob(args struct{ ID gql.ID }) *PlaylistJobResolver {
	job := &PlaylistJob{id: args.ID, url: ""}

	return &PlaylistJobResolver{job}
}
