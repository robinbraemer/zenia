package server

import (
	"context"
	"errors"
	"zenia/pkg/acl"
)

// Server is a Zenia ACL server that is organized in a cluster.
// It is an Access Control List server that performs authorization
// checks based on stored permissions specified by clients.
//
// Requests arrive at any server in a cluster and that server fans
// out the work to other servers in the cluster as necessary.
// Those servers may in turn contact other servers to compute
// intermediate results. The initial server gathers the final
// result and returns it to the client.
//
// This system helps establish consistent semantics and user experience
// across applications to interoperate and coordinate access
// control. Common infrastructure can be built on top of this
// unified globally scalable access control system.
type Server struct {
	Storage Store
}

var (
	// ErrStorageItemNotFound is the error returned by Store implementations
	// to indicate that a requested item was not found.
	ErrStorageItemNotFound = errors.New("store: item not found")
)

// Store provides the persistent storage abstraction to the ACL servers,
// which read and write those databases in the course of
// responding to client requests.
//
// There is for each a database for:
//	- relation tuples for each (client specified) namespace
//	- all namespace configurations
//	- changelog shared across all namespaces
//
type Store interface {
	// Gets all namespaces with it's config.
	GetNamespaces(context.Context) ([]acl.Namespace, error)
}
