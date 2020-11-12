#!/bin/bash

# Generate protobuf & grpc API
buf generate

# Generate gapic client
rm -r client/authz/apiv1 2> /dev/null
buf generate \
	--config '{"version":"v1beta1","build":{"roots":["api"],"excludes":["api/zenia/node"]}}' \
	--template '{"version":"v1beta1","plugins":[{"name":"go_gapic","out":"client","opt":
	"go-gapic-package=authz/apiv1;authz"}]}'
