package memory

import (
	"context"
	"fmt"
	"strings"
	"time"
	"zenia/pkg/acl"
)

// New returns a new in-memory server.Store.
// It can only be used to run a single server
// or local cluster in the same process for testing!
func New() *Memory {
	return &Memory{
		tuplesByNamespace: map[string][]item{},
		namespaces:        namespaces{},
	}
}

type Memory struct {
	tuplesByNamespace map[string][]item
	namespaces        namespaces
}
type namespaces []acl.Namespace

func (d *Memory) GetNamespaces(ctx context.Context) ([]acl.Namespace, error) {
	return d.namespaces, nil
}

func (d *Memory) SaveNamespace(
	_ context.Context, namespace acl.Namespace,
) error {
	for i, ns := range d.namespaces {
		if ns.Name == namespace.Name {
			d.namespaces[i] = namespace // replace existing
			return nil
		}
	}
	d.namespaces = append(d.namespaces, namespace)
	return nil
}

// Implement this interface
//var _ server.Store = (*Memory)(nil)

// By: shardID, objectID, relation, user, commit timestamp
type item struct {
	shardID string
	acl.RelationTuple
	commitTime time.Time
}

func (d *Memory) Exists(
	_ context.Context, object acl.Object,
	relation string, userID string,
) (ok bool, err error) {
	for _, i := range d.tuplesByNamespace[object.Namespace] {
		if !(strings.EqualFold(i.Relation, relation) &&
			strings.EqualFold(i.Object.ID, object.ID)) {
			continue
		}

		if len(i.User.ID) != 0 &&
			strings.EqualFold(i.User.ID, userID) {
			return true, nil
		}

		ok2, err := d.Exists(context.TODO(), acl.Object{
			Namespace: i.User.UserSet.Object.Namespace,
			ID:        i.User.UserSet.Object.ID,
		}, i.Relation, userID)
		if err != nil {
			return false, fmt.Errorf(
				"error resolving checking relationship: %w", err)
		}
		if ok2 {
			return true, nil
		}
	}
	return
}

func (d *Memory) UserSets(
	_ context.Context, object acl.Object, relation string,
) (set []acl.UserSet, err error) {
	t, ok := d.tuplesByNamespace[object.Namespace]
	if !ok {
		return
	}
	for _, i := range t {
		// optionally filter by relation
		if len(relation) != 0 &&
			!strings.EqualFold(i.Relation, relation) {
			continue
		}
		set = append(set, i.User.UserSet)
	}
	return
}

func (d *Memory) Save(_ context.Context, tuple acl.RelationTuple) error {
	d.tuplesByNamespace[tuple.Object.Namespace] = append(
		d.tuplesByNamespace[tuple.Object.Namespace], item{
			shardID:       "",
			RelationTuple: tuple,
			commitTime:    time.Now().UTC(),
		})
	return nil
}

//func (d *Memory) GetRelationTuples(_ context.Context,
//	ns string, oid string, relation string,
//) (results []acl.RelationTuple, err error) {
//	rt, ok := d.tuplesByNamespace[ns]
//	if !ok {
//		return nil, server.ErrStorageItemNotFound
//	}
//	for _, i := range rt {
//		// Does item in Memory match request?
//		if !strings.EqualFold(string(i.objectID), string(oid)) {
//			continue
//		}
//		// relation is optional
//		if len(relation) != 0 && !strings.EqualFold(string(i.relation), string(relation)) {
//			continue
//		}
//
//		// Add item to results
//		results = append(results, acl.RelationTuple{
//			Object: acl.Object{
//				Namespace: ns,
//				ID:  i.objectID,
//			},
//			Relation: relation,
//			User:     i.user,
//		})
//	}
//	return
//}
