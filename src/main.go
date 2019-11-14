package main

import (
	"log" 
	"net/http"

	// "time"

	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jonathanbruce/pnin/src/graphql"
)

func respondWithSchema(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(graphql.SchemaString))
}

func main() {
	// var (
	// 	addr              = ":8000"
	// 	readHeaderTimeout = 1 * time.Second
	// 	writeTimeout      = 10 * time.Second
	// 	idletimeout       = 90 * time.Second
	// 	maxHeaderBytes    = http.DefaultMaxHeaderBytes
	// )

	schema := graphql.GetSchema()
	var addr = "3000"

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	printWelcome(addr)

	http.Handle("/graphql", &relay.Handler{Schema: schema})
	http.HandleFunc("/schema.graphql", respondWithSchema)
	log.Fatal(http.ListenAndServe(":"+addr, nil))
}
