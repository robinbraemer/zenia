package server

import (
	"context"
	"errors"
	"github.com/robinbraemer/zenia/pkg/acl"
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
	store Store
	// <namespace, <relation, V>>
	userSetRewrites map[string]map[string]*acl.UserSetRewrite
}

// LoadNamespaces reloads all namespaces and it's config from the store.
func (s *Server) LoadNamespaces(ctx context.Context) error {
	nsc, err := s.store.GetNamespaces(ctx)
	if err != nil {
		return err
	}

	if s.userSetRewrites == nil {
		s.userSetRewrites = map[string]map[string]*acl.UserSetRewrite{}
	}
	setUSR := func(namespace, relation string, rule *acl.UserSetRewrite) {
		n, ok := s.userSetRewrites[namespace]
		if !ok {
			n = map[string]*acl.UserSetRewrite{}
			s.userSetRewrites[namespace] = n
		}
		n[relation] = rule
	}
	for _, ns := range nsc {
		for _, rel := range ns.Relations {
			setUSR(ns.Name, rel.Name, &rel.UserSetRewrite)
		}
	}
	return nil
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
	TupleStore
	NamespaceStore
}

type Service interface {
	Read(ctx context.Context)
}

// TupleStore specifies the store needed to store and retrieve ACLs.
// An implementation must be very fast, especially for reads.
// ACL servers reads and writes this store in the course of
// responding to client requests.
//
// There is for each a database for:
//	- relation tuples for each (client specified) namespace
//	- all namespace configurations
//	- changelog shared across all namespaces
//
type TupleStore interface {
	// Checks whether the exact tuple exists in the store.
	Exists(ctx context.Context, object acl.Object, relation string, userID string) (bool, error)
	// Gets the users filtered by an object and
	// optionally it's corresponding relation.
	UserSets(ctx context.Context, object acl.Object, relation string) ([]acl.UserSet, error)
	// Saves writes a new relation tuple entry to the store.
	Save(ctx context.Context, tuple acl.RelationTuple) error
}

type NamespaceStore interface {
	// Gets all namespaces with it's config.
	GetNamespaces(context.Context) ([]acl.Namespace, error)
	// Saves a namespace with it's config.
	SaveNamespace(context.Context, acl.Namespace) error
}
