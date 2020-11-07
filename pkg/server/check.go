package server

import (
	"context"
	"fmt"
	"zenia/pkg/acl"
)

// CheckRequest is a request to a Server.Check and
// specifies one or multiple tuplesets and an optional
// zookie.
type CheckRequest struct {
	Object   acl.Object
	Relation string
	UserID   string

	Zookie []byte // Optional
}

type CheckResponse struct {
	Exists bool
}

// Check answers the question of
// "Does user U have relation R to object O?".
//
// Clients read relation tuples to display
// ACLs or group membership to users, or to
// prepare for a subsequent write.
func (s *Server) Check(ctx context.Context, r *CheckRequest) (*CheckResponse, error) {
	/*

		(consider new table per namespace)
		namespace | object_id | relation | user_id | userset_namespace | userset_object_id | userset_relation
		----------+-----------+----------+---------+-------------------+-------------------+-------------------
		doc       | readme    | owner    | 10      | NULL              | NULL              | NULL
		group	  | eng       | member   | 11      | NULL              | NULL              | NULL
		doc       | readme    | viewer   | NULL    | group             | eng               | member
		doc       | readme    | parent   | NULL    | folder            | A                 | ...

		User 10 is an owner of doc:readme
		User 11 is a member of group:eng
		Members of group:eng are viewers of doc:readme
		doc:readme is in folder:A

	*/

	// Check store directly has this tuple (No inheritance).
	ok, err := s.store.Exists(ctx, r.Object, r.Relation, r.UserID)
	if err != nil {
		return nil, fmt.Errorf(
			"error checking store has tuple (%s->%s UserID: %s): %w",
			r.Object, r.Relation, r.UserID, err)
	}
	if ok {
		return &CheckResponse{Exists: true}, nil
	}

	userSets, err := s.store.UserSets(ctx, r.Object, r.Relation)
	if err != nil {
		return nil, fmt.Errorf("error getting usersets for %s -> %s: %w",
			r.Object, r.Relation, err)
	}

	rule := s.getUserSetRewrite(r.Object.Namespace, r.Relation)
	if rule != nil {
		userSets = append(userSets, computedUserSet(rule, r.Object)...)
	}

	for _, set := range userSets {
		// TODO in parallel... see paper
		res, err := s.Check(ctx, &CheckRequest{
			Object:   set.Object,
			Relation: set.Relation,
			UserID:   r.UserID,
			Zookie:   r.Zookie,
		})
		if err != nil {
			return nil, err
		}
		if res.Exists {
			return &CheckResponse{Exists: true}, nil
		}
	}

	return &CheckResponse{Exists: ok}, nil
}

func computedUserSet(rule *acl.UserSetRewrite, object acl.Object) (set []acl.UserSet) {
	if rule == nil {
		return
	}
	for _, u := range rule.Union {
		if u.ComputedUserSet.Relation == "" {
			continue
		}
		set = append(set, acl.UserSet{
			Object:   object,
			Relation: u.ComputedUserSet.Relation,
		})
	}
	return
}

// returns cached userset rewrite rule
func (s *Server) getUserSetRewrite(namespace, relation string) *acl.UserSetRewrite {
	m, ok := s.userSetRewrites[namespace]
	if !ok {
		return nil
	}
	return m[relation]
}
