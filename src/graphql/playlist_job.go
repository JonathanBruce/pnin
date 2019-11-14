package graphql

import gql "github.com/graph-gophers/graphql-go"

var PlaylistDef = `
type PlaylistJob {
	id: ID!
	playlistURL: String!
}
`

type PlaylistJob struct {
	id  gql.ID
	url string
}

type PlaylistJobResolver struct {
	job *PlaylistJob
}

// PlaylistURL Returns the playlist job's playlist url string
func (p *PlaylistJobResolver) PlaylistURL() string {
	return p.job.url
}

// ID Returns the pnin job id
func (p *PlaylistJobResolver) ID() gql.ID {
	return p.job.id
}

// PlaylistJob returns a playlist job for a given playlistId
func (p *PlaylistJobResolver) PlaylistJob(playlistID gql.ID) *PlaylistJob {
	return &PlaylistJob{id: playlistID, url: ""}
}
