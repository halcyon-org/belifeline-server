//go:build tools
// +build tools

package main

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/google/wire/cmd/wire"
)
