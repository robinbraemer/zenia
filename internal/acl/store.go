package acl

import (
	"encoding/json"
	"github.com/robinbraemer/zenia/api/zenia/authz/v1"
	"strings"
)

type Store interface {
}

type RelationTuple struct {
	*authz.RelationTuple
}

func (r *RelationTuple) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name string `json:"name"`
	}{
		Name: strings.ToUpper(p.Name),
	})
}

var _ json.Marshaler = (*RelationTuple)(nil)
