package server

import (
	"context"
)

type (
	// ReadRequest is a request to a Server.Read and
	// specifies one or multiple tuplesets and an optional
	// zookie.
	ReadRequest struct
		// Tuples<RelationTuple-Key, []>
		Tuples map[TupleKey][]TupleKey
		Zookie []byte // Optional
	}
	ReadResponse struct{}
)

type TupleKey string

// Read
//
// Clients read relation tuples to display
// ACLs or group membership to users, or to
// prepare for a subsequent write.
func (s *Server) Read(
	ctx context.Context, in *ReadRequest,
) (*ReadResponse, error) {

}
