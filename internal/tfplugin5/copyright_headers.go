// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-1.1

//go:build tools
// +build tools

package tfplugin5

import (
	_ "github.com/hashicorp/copywrite"
)

//go:generate go run github.com/hashicorp/copywrite headers
