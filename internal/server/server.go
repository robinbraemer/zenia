package server

import (
	admin "github.com/robinbraemer/zenia/api/zenia/authz/admin/v1"
	api "github.com/robinbraemer/zenia/api/zenia/authz/v1"
	node "github.com/robinbraemer/zenia/api/zenia/node/v1"
)

// Server is a Zenia ACL server that is organized in a cluster.
// It is an Access Control List server that performs authorization
// checks based on stored permissions specified by clients.
//
// Requests arrive at any server in a cluster and that server fans
// out the work to other servers in the cluster as necessary.
// Those servers may in turn contact other servers to compute
// intermediate results. The initial server gathers the final
// result and returns it to the client.
//
// This system helps establish consistent semantics and user experience
// across applications to interoperate and coordinate access
// control. Common infrastructure can be built on top of this
// unified globally scalable access control system.
type Server struct {
}

var _ api.AuthorizationServiceServer = (*Server)(nil)
var _ node.NodeServiceServer = (*Server)(nil)
var _ admin.AdminServiceServer = (*Server)(nil)
