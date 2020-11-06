package acl

type UserSetExpression struct {
	Sub []*UserSetExpression
}

// UserSetRewriteRule is defined per relation in a namespace
// and specifies a function that takes and ObjectID and returns
// a UserSetExpressionTree.
type UserSetRewriteRule interface {
	Rewrite(objectID string) *UserSetExpressionTree
}

// UserSetExpressionTree is an expression tree returned by
// a UserSetRewriteRule function.
type UserSetExpressionTree UserSetExpressionTreeNode

// UserSetExpressionTreeNode is a leaf node in a
// UserSetExpressionTree and has only one of it's fields
// set.
type UserSetExpressionTreeNode struct {
	// All users from stored RelationTuples for the
	// <object>#<relation> pair, including indirect
	// ACLs referenced by UserSets from the tuples.
	//
	// The default behaviour when no rewrite rule is specified.
	This struct {
		UserSet                              // Users in this direct <object>#<relation> pair.
		Refs    []*UserSetExpressionTreeNode // TODO ACLs referenced by user sets.
	}
	// The new UserSet computed from the input ObjectID.
	// E.g. allows the userset expression for a "viewer"
	// relation to refer to the "editor" relation user set
	// on the same object, thus offering an ACL inheritance
	// capability between relations.
	ComputedUserSet UserSet
	// Refer to TupleToUserSet interface.
	TupleToUserSet TupleToUserSet
}

// TupleToUserSet computes a RelationTupleSet from the input object,
// fetches RelationTuples matching the RelationTupleSet,
// and computes a userset from every fetched relation tuple.
// This flexible primitive allows our clients to express complex
// policies such as “look up the parent folder of the document
// and inherit its viewers”.
type TupleToUserSet interface {
	//Fetch(objectID string) RelationTupleSet
}
