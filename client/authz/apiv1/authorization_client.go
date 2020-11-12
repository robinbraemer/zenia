// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package authz

import (
	"context"
	"math"

	gax "github.com/googleapis/gax-go/v2"
	authzpb "github.com/robinbraemer/zenia/api/authz/zenia/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newAuthorizationClientHook clientHook

// AuthorizationCallOptions contains the retry settings for each method of AuthorizationClient.
type AuthorizationCallOptions struct {
	Check []gax.CallOption
}

func defaultAuthorizationClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("authorization.exampleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("authorization.exampleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAuthorizationCallOptions() *AuthorizationCallOptions {
	return &AuthorizationCallOptions{
		Check: []gax.CallOption{},
	}
}

// AuthorizationClient is a client for interacting with .
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type AuthorizationClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	authorizationClient authzpb.AuthorizationServiceClient

	// The call options for this service.
	CallOptions *AuthorizationCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAuthorizationClient creates a new authorization service client.
//
func NewAuthorizationClient(ctx context.Context, opts ...option.ClientOption) (*AuthorizationClient, error) {
	clientOpts := defaultAuthorizationClientOptions()

	if newAuthorizationClientHook != nil {
		hookOpts, err := newAuthorizationClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &AuthorizationClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultAuthorizationCallOptions(),

		authorizationClient: authzpb.NewAuthorizationServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AuthorizationClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AuthorizationClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AuthorizationClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

func (c *AuthorizationClient) Check(ctx context.Context, req *authzpb.CheckRequest, opts ...gax.CallOption) (*authzpb.CheckResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.Check[0:len(c.CallOptions.Check):len(c.CallOptions.Check)], opts...)
	var resp *authzpb.CheckResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.authorizationClient.Check(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
