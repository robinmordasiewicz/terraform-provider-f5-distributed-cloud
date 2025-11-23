//go:build tools

// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

// Package tools tracks tool dependencies for this module.
// This file ensures that tool dependencies are tracked in go.mod
// and can be installed via `go install`.
package tools

import (
	// Documentation generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)
