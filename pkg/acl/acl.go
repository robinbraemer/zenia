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

// ParentRelation is used to generally
// refer to parent relation tuples.
const ParentRelation = "parent"

// Object is the object mapped by a RelationTuple.
type Object struct {
	Namespace string // The namespace of the Object.
	ID        string // The object ID.
}

// UserSet refers to a set of users.
type UserSet struct {
	Object   Object
	Relation string
}

// User is the object mapped by a RelationTuple
// and contains either a user id or a set of users.
type User struct {
	ID      string  // A user id.
	UserSet UserSet // A user set.
}

// RelationTuple is an ACL entry that relates an Object to a User.
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
	Name           string           `yaml:"name"` // The name of the relation.
	RelationConfig `yaml:",inline"` // The configuration.
}

// RelationConfig is a configuration for an ACL relation.
type RelationConfig struct {
	// Config for rewriting user set relations.
	UserSetRewrite UserSetRewrite `yaml:"userset_rewrite"`
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
	Name            string           `yaml:"name"` // The name.
	NamespaceConfig `yaml:",inline"` // The configuration.
}

// NamespaceConfig is the configuration for an ACL namespace
// to separate tenants. Clients must first configure one
// before storing relation tuples in it.
//
// This config specifies its relations and store parameters.
type NamespaceConfig struct {
	// The relations by name configured in the namespace.
	Relations []Relation `yaml:"relations"`
	// Storage settings for the specified Relations
	// in the namespace.
	Storage NamespaceStorageSettings `yaml:"storage"`
}

// NamespaceStorageSettings are the store settings
// for Relations contained in a Namespace.
type NamespaceStorageSettings struct {
	// Sharding settings for the tuples in the namespace.
	Sharding struct {
		// Defaults to "$OBJECT_ID",
		// meaning shard id == object id.
		ComputedIDExpression string `yaml:"computed_id_expression"`
	} `yaml:"sharding,omitempty"`
	// The encoding to optimize store of integer,
	// string and other ObjectID formats.
	//ObjectIDEncoding interface{}
}
