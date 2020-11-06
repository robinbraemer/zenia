package server

import (
	"context"
)

// ReadRequest is a request to a Server.Read and
// specifies one or multiple tuplesets and an optional
// zookie.
type ReadRequest struct {
	// Tuples<RelationTupleKey, []TupleKey>
	Tuples map[string][]string
	Zookie []byte // Optional
}

type ReadResponse struct{}

// Read
//
// Clients read relation tuples to display
// ACLs or group membership to users, or to
// prepare for a subsequent write.
func (s *Server) Read(ctx context.Context, in *ReadRequest) (*ReadResponse, error) {
	return nil, nil
}
