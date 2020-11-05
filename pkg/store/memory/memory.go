package memory

import (
	"context"
	"strings"
	"time"
	"zenia/pkg/acl"
	"zenia/pkg/server"
)

// New returns a new in-memory server.Store.
// It can only be used to run a single server
// or local cluster in the same process for testing!
func New() server.Store {
	return &db{
		relationTuplesByNamespace: map[acl.NamespaceName]relationTuples{},
		namespaces:                namespaces{},
	}
}

type db struct {
	relationTuplesByNamespace map[acl.NamespaceName]relationTuples
	namespaces                namespaces
}

// Implement this interface
var _ server.Store = (*db)(nil)

// relation tuples database
type (
	relationTuples map[relationTuplePrimaryKey]relationTupleItem
	// By: shardID, objectID, relation, user, commit timestamp
	relationTuplePrimaryKey string
	relationTupleItem       struct {
		shardID    string
		objectID   acl.ObjectID
		relation   acl.RelationName
		user       acl.User
		commitTime time.Time
	}
)

func (d *db) GetRelationTuples(_ context.Context,
	ns acl.NamespaceName, oid acl.ObjectID, relation acl.RelationName,
) (results []acl.RelationTuple, err error) {
	rt, ok := d.relationTuplesByNamespace[ns]
	if !ok {
		return nil, server.ErrStorageItemNotFound
	}
	for _, i := range rt {
		// Does item in db match request?
		if !strings.EqualFold(string(i.objectID), string(oid)) {
			continue
		}
		// relation is optional
		if len(relation) != 0 && !strings.EqualFold(string(i.relation), string(relation)) {
			continue
		}

		// Add item to results
		results = append(results, acl.RelationTuple{
			Object: acl.Object{
				Namespace: ns,
				ObjectID:  i.objectID,
			},
			Relation: relation,
			User:     i.user,
		})
	}
	return
}

// namespaces with its config database
type (
	namespaces []acl.Namespace
	//namespaceChangelog map[acl.NamespaceName] TODO
)

func (d *db) GetNamespaces(ctx context.Context) ([]acl.Namespace, error) {
	return d.namespaces, nil
}
