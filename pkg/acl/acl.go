/*
The acl package provides the implementation of ACL data model.
ACLs are defined around tuples, instead of per-object ACLs,
that instead allows to unify the concepts of ACLs and groups
and support efficient reads and incremental updates.

Those relation tuples have efficient binary encodings,
but are representable (in code comments) as the following
text notations:

	RelationTuple:
		<tuple> = <object>#<relation>@<user>
		<object> = <namespace>:<object_id>
		<user> = <user_id> | <userset>
		<userset> = <object>#<relation>

where:

	<namespace> and <relation> are predefined in client configurations
	<object_id> is a string
	<user_id> is an integer

The primary keys required to identify a relation tuple are:

	<namespace>,
	<object_id>,
	<relation> and
	<user>
*/
package acl

import "context"

// ParentRelation is used to generally
// refer to parent relation tuples.
const ParentRelation = "parent"

// Object is the object mapped by a RelationTuple.
type Object struct {
	Namespace string // The namespace of the Object.
	ID        string // The object ID.
}

// UserSet is a set of user ids.
type UserSet map[string]struct{}

// User is the object mapped by a RelationTuple
// and contains either a user id or a set of users.
type User struct {
	ID      string  // A user id.
	UserSet UserSet // A user set.
}

// RelationTuple is an ACL entry that relates an
// Object to a User (which maybe a UserSet for inheritance).
//
// While this tuple reflects a relationship between Objects
// and Users, they do not completely define the effective ACLs.
type RelationTuple struct {
	Object   Object
	Relation string
	User     User
}

// Relation is a relation in a Namespace in the ACL system
// that ties Object and User together in a RelationTuple.
type Relation struct {
	Name           string // The name of the relation.
	RelationConfig        // The configuration.
}

// RelationConfig is a configuration for an ACL relation.
type RelationConfig struct {
	// Config for rewriting user set relations.
	UserSetRewrite UserSetRewrite
}

// UserSetRewrite offers clients highly configurable
// behaviour for authorization checks.
type UserSetRewrite struct {
	Union []struct {
		ComputedUserSet struct {
			Relation string `yaml:"relation"`
		} `yaml:"computed_userset,omitempty"`
		TupleToUserSet struct {
			TupleSet struct {
				Relation string `yaml:"relation"`
			} `yaml:"tupleset"`
			ComputedUserSet struct {
				Object   string `yaml:"object"`
				Relation string `yaml:"relation"`
			} `yaml:"computed_userset"`
		} `yaml:"tuple_to_userset,omitempty"`
	} `yaml:"union"`
}

// Namespace is a namespace in the ACL system
// that contains multiple Relations.
type Namespace struct {
	Name            string // The name.
	NamespaceConfig        // The configuration.
}

// NamespaceConfig is the configuration for an ACL namespace
// to separate tenants. Clients must first configure one
// before storing relation tuples in it.
//
// This config specifies its relations and store parameters.
type NamespaceConfig struct {
	// The relations by name configured in the namespace.
	Relations map[string]*Relation
	// Storage settings for the specified Relations
	// in the namespace.
	Storage NamespaceStorageSettings
}

// NamespaceStorageSettings are the store settings
// for Relations contained in a Namespace.
type NamespaceStorageSettings struct {
	// Sharding settings for the tuples in the namespace.
	Sharding struct{}
	// The encoding to optimize store of integer,
	// string and other ObjectID formats.
	ObjectIDEncoding interface{}
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
	Exists(ctx context.Context, object Object, relation string, userID string) (bool, error)
	// Gets the users filtered by an object and
	// optionally it's corresponding relation.
	UserSets(ctx context.Context, object Object, relation string) ([]UserSet, error)
	// Saves writes a new relation tuple entry to the store.
	Save(ctx context.Context, tuple RelationTuple) error
}
