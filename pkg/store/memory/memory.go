package memory

import (
	"context"
	"fmt"
	"strings"
	"time"
	"zenia/pkg/acl"
)

// New returns a new in-memory server.Store.
// It is highly unoptimized, can only run in a
// single server or local cluster in the same process
// and should only be used for local testing!
func New() *Memory {
	return &Memory{
		tuples:     map[string]item{},
		namespaces: namespaces{},
	}
}

type Memory struct {
	tuples     map[string]item
	namespaces namespaces
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

type item struct {
	acl.RelationTuple
	commitTime time.Time
}

func (d *Memory) Exists(
	ctx context.Context, object acl.Object,
	relation string, userID string,
) (ok bool, err error) {

	type expandItem struct {
		object   acl.Object
		relation string
	}

	var expand []expandItem
	for _, i := range d.tuples {
		if !(strings.EqualFold(i.Object.Namespace, object.Namespace) &&
			strings.EqualFold(i.Relation, relation) &&
			strings.EqualFold(i.Object.ID, object.ID)) {
			continue
		}

		if i.User.ID != "" &&
			strings.EqualFold(i.User.ID, userID) {
			return true, nil
		}

		expand = append(expand, expandItem{
			object: acl.Object{
				Namespace: i.User.UserSet.Object.Namespace,
				ID:        i.User.UserSet.Object.ID,
			},
			relation: i.Relation,
		})
	}
	for _, e := range expand {
		ok, err := d.Exists(ctx, e.object, e.relation, userID)
		if err != nil {
			return false, fmt.Errorf(
				"error checking expanded relationship: %w", err)
		}
		if ok {
			return true, nil
		}
	}
	return
}

func (d *Memory) UserSets(
	_ context.Context, object acl.Object, relation string,
) (set []acl.UserSet, err error) {
	for _, t := range d.tuples {
		if !strings.EqualFold(t.Object.Namespace, object.Namespace) {
			continue
		}
		if !strings.EqualFold(t.Object.ID, object.ID) {
			continue
		}
		// optionally filter by relation
		if relation != "" &&
			!strings.EqualFold(t.Relation, relation) {
			continue
		}
		set = append(set, t.User.UserSet)
	}
	return
}

func (d *Memory) Save(_ context.Context, tuple acl.RelationTuple) error {
	i := item{
		RelationTuple: tuple,
		commitTime:    time.Now().UTC(),
	}
	d.tuples[itemKey(i)] = i
	return nil
}

func itemKey(i item) string {
	return i.commitTime.String() + " " +
		objectString(i.Object) + "#" +
		i.Relation + "@" +
		userString(i.User)
}

func objectString(o acl.Object) string {
	return o.Namespace + ":" + o.ID
}

func userString(u acl.User) string {
	if u.ID == "" {
		return u.UserSet.Object.Namespace + "#" + u.UserSet.Object.ID
	}
	return u.ID
}

//func (d *Memory) GetRelationTuples(_ context.Context,
//	ns string, oid string, relation string,
//) (results []acl.RelationTuple, err error) {
//	rt, ok := d.tuples[ns]
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
