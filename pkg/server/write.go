package server

import "context"

type (
	WriteRequest struct {
	}
	WriteResponse struct{}
)

// Write
//
// Clients may modify a single relation tuple to
// add or remove an ACL. They may also modify all
// tuples related to an object via a read-modify-write
// process with optimistic concurrency control that uses
// a Server.Read followed by a Server.Write.
//
//	1. Read all relation tuples of an object,
//	   including a per-object "lock" tuple.
//  2. Generate the tuples to write or delete.
// 	   Send the writes, along with a touch on the
//	   lock tuple, to Zenia, with the condition that
//	   the writes will be committed only if the lock
//	   tuple has not been modified since the read.
//  3. If the write condition is not met, go back to step 1.
//
// The lock tuple is just a regular relation tuple used by
// clients to detect write races.
//
func (s *Server) Write(
	ctx context.Context, in *WriteRequest,
) (*WriteResponse, error) {
	return nil, nil
}
