// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure F5XCProvider satisfies various provider interfaces.
var _ provider.Provider = &F5XCProvider{}

// F5XCProvider defines the provider implementation.
type F5XCProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// New returns a new provider instance.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &F5XCProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *F5XCProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "f5xc"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *F5XCProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Terraform provider for F5 Distributed Cloud (F5 XC) services.",
		Attributes:  map[string]schema.Attribute{},
	}
}

// Configure prepares an F5 XC API client for data sources and resources.
func (p *F5XCProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// TODO: Implement provider configuration in Phase 2/3
}

// Resources defines the resources implemented in the provider.
func (p *F5XCProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

// DataSources defines the data sources implemented in the provider.
func (p *F5XCProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
