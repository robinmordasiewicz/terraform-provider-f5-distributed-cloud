// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package timeouts

import (
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Default timeout values for F5 XC operations.
const (
	// DefaultCreate is the default timeout for create operations.
	DefaultCreate = 30 * time.Minute

	// DefaultRead is the default timeout for read operations.
	DefaultRead = 5 * time.Minute

	// DefaultUpdate is the default timeout for update operations.
	DefaultUpdate = 30 * time.Minute

	// DefaultDelete is the default timeout for delete operations.
	DefaultDelete = 30 * time.Minute

	// SiteCreate is the timeout for site creation (longer due to provisioning).
	SiteCreate = 60 * time.Minute

	// SiteDelete is the timeout for site deletion.
	SiteDelete = 60 * time.Minute
)

// BlockAll returns a timeouts block with all CRUD timeouts enabled.
func BlockAll() schema.Block {
	return timeouts.Block(
		nil,
		timeouts.Opts{
			Create: true,
			Read:   true,
			Update: true,
			Delete: true,
		},
	)
}

// BlockCreateUpdateDelete returns a timeouts block without read timeout.
func BlockCreateUpdateDelete() schema.Block {
	return timeouts.Block(
		nil,
		timeouts.Opts{
			Create: true,
			Update: true,
			Delete: true,
		},
	)
}

// BlockSite returns a timeouts block with longer timeouts for site resources.
func BlockSite() schema.Block {
	return timeouts.Block(
		nil,
		timeouts.Opts{
			Create: true,
			Read:   true,
			Update: true,
			Delete: true,
		},
	)
}
