package server

import (
	"context"
	"fmt"
	"github.com/robinbraemer/zenia/pkg/acl"
	crdbstore "github.com/robinbraemer/zenia/pkg/store/crdb"
	"github.com/robinbraemer/zenia/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
	"time"
)

// Tip: Do read from right to left.
var tuples = []string{
	"group:sales#member@11",
	"doc:readme#owner@10",
	"doc:readme#viewer@group:sales#member",
	"doc:readme#parent@folder:ABC#...", // doc:readme is in folder:ABC
	// "#..." represents a relation that does not affect the semantics of the tuple

	//"doc:readme#viewer@13",
	//
	//"group:all#member@group:devops#member",
	//"doc:readme#viewer@group:all#member",
	//
	//
	//"group:sales#member@12",
	//"folder:A#viewer@group:sales#member",
}

//var videoTuples = []string{
//"video:holmes#viewer@1",
//"group:eng#member@3",
//"video:holmes#parent@channel:audiobooks#...",
//"channel:audiobooks#editor@group:eng#member",
//"channel:audiobooks#viewer@2",
//}

var checks = []struct {
	name     string
	expected bool
	given    string
}{
	{
		name:     "User 11 is member of the sales group",
		expected: true,
		given:    "group:sales#member@11",
	},
	{ // Given through "doc:readme#viewer@group:sales#member"
		name:     "Members of sales group are viewers of doc:readme",
		expected: true,
		given:    "doc:readme#viewer@11",
	},

	// Test "case sensitivity".
	{
		name:     "case sensitivity: User 11 is not member of the sales GROUP",
		expected: false,
		given:    "GROUP:sales#member@11",
	},
	{
		name:     "case sensitivity: User 11 is not member of the SALES group",
		expected: false,
		given:    "group:SALES#member@11",
	},
	{
		name:     "case sensitivity: User 11 is not MEMBER of the sales group",
		expected: false,
		given:    "group:sales#MEMBER@11",
	},

	//{
	//	name:     "not found relation on nonexistent user",
	//	expected: false,
	//	given:    "doc:readme#viewer@15",
	//},
	//{
	//	name:     "not found relation",
	//	expected: false,
	//	given:    "doc:readme#viewer@10",
	//},
	//{
	//	name:     "not found relation",
	//	expected: false,
	//	given:    "doc:readme#member@12",
	//},
	//{
	//	name:     "relation by group found",
	//	expected: true,
	//	given:    "doc:readme#viewer@11",
	//},
	//
	//{
	//	name:     "not found relation",
	//	expected: true,
	//	given:    "doc:readme#viewer@13",
	//},
	//{
	//	name:     "found by parent -> editor -> group -> member",
	//	expected: true,
	//	given:    "video:holmes#viewer@3",
	//},
	//{
	//	name:     "found by parent -> viewer",
	//	expected: true,
	//	given:    "video:holmes#viewer@2",
	//},
}

var store *crdbstore.Crdb

func TestMain(t *testing.M) {
	if err := testMain(t); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func testMain(t *testing.M) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	store, err = crdbstore.NewTestStore(ctx, "../store/crdb/all.sql")
	if err != nil {
		return fmt.Errorf("error new cockroachdb store: %v", err)
	}

	// Run tests
	if exist := t.Run(); exist != 0 {
		return fmt.Errorf("test returned zero-exit code = %d", exist)
	}
	return nil
}

func TestChecks(t *testing.T) {
	// Create & init store
	require.NoError(t, initStore(store))

	// Create and init auth server
	auth := &Server{store: store}
	err := auth.LoadNamespaces(context.TODO())
	require.NoError(t, err)

	for _, tt := range checks {
		t.Run(tt.name, func(t *testing.T) {
			tuple := parse(tt.given)
			res, err := auth.Check(context.TODO(), &CheckRequest{
				Object:   tuple.Object,
				Relation: tuple.Relation,
				UserID:   tuple.User.ID,
			})
			if !assert.NoError(t, err) {
				return
			}
			actual := res.Exists
			assert.Equal(t, tt.expected, actual, "%s", tt.given)
		})
	}
}

func parse(tuple string) (t acl.RelationTuple) {
	parts := strings.Split(tuple, "@")

	objectRelation := strings.Split(parts[0], "#")
	nsObjectID := strings.Split(objectRelation[0], ":")

	t = acl.RelationTuple{
		Object: acl.Object{
			Namespace: nsObjectID[0],
			ID:        nsObjectID[1],
		},
		Relation: objectRelation[1],
	}

	parts = strings.Split(parts[1], "#")
	if len(parts) == 1 {
		t.User.ID = parts[0]
	} else {
		objectRelation = parts
		nsObjectID = strings.Split(objectRelation[0], ":")
		t.User.UserSet = acl.UserSet{
			Object: acl.Object{
				Namespace: nsObjectID[0],
				ID:        nsObjectID[1],
			},
			Relation: objectRelation[1],
		}
	}
	return
}

func initStore(store Store) error {
	nsc, err := testdata.LoadNamespaceConfigs("../../testdata")
	if err != nil {
		return err
	}
	for _, ns := range nsc {
		err = store.SaveNamespace(context.TODO(), ns)
		if err != nil {
			return fmt.Errorf(
				"error saving namespace config %s: %w",
				ns.Name, err)
		}
	}

	for _, unparsed := range tuples {
		tuple := parse(unparsed)
		if err := store.Save(context.TODO(), tuple); err != nil {
			return fmt.Errorf("error saving tuple %q: %w", unparsed, err)
		}
	}
	return nil
}
