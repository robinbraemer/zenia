package server

import (
	"context"
	admin "github.com/robinbraemer/zenia/api/zenia/authz/admin/v1"
)

func (s *Server) ApplyNamespace(ctx context.Context, in *admin.ApplyNamespaceRequest) (*admin.ApplyNamespaceResponse, error) {
	panic("implement me")
}
