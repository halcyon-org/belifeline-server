//go:build tools
// +build tools

package main

import (
	"connectrpc.com/connect/cmd/protoc-gen-connect-go"
	"github.com/bufbuild/buf/cmd/buf"
	"github.com/google/wire/cmd/wire"
	"google.golang.org/protobuf/cmd/protoc-gen-go"
)
