// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package api_credential

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &APICredentialDataSource{}

func NewAPICredentialDataSource() datasource.DataSource {
	return &APICredentialDataSource{}
}

type APICredentialDataSource struct {
	client *client.Client
}

type APICredentialDataSourceModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Namespace      types.String `tfsdk:"namespace"`
	Description    types.String `tfsdk:"description"`
	CredentialType types.String `tfsdk:"credential_type"`
}

type APIAPICredential struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		CredentialType string `json:"credential_type,omitempty"`
	} `json:"spec"`
}

func (d *APICredentialDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_credential"
}

func (d *APICredentialDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud API Credential.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_api_credential`" + ` data source retrieves information about an existing API credential.

## Example Usage

` + "```hcl" + `
data "f5_distributed_cloud_api_credential" "example" {
  name      = "my-api-credential"
  namespace = "system"
}

output "credential_type" {
  value = data.f5_distributed_cloud_api_credential.example.credential_type
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the API credential.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the API credential to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the API credential is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the API credential.",
				Computed:    true,
			},
			"credential_type": schema.StringAttribute{
				Description: "Type of the API credential.",
				Computed:    true,
			},
		},
	}
}

func (d *APICredentialDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData))
		return
	}
	d.client = c
}

func (d *APICredentialDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data APICredentialDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIAPICredential
	path := fmt.Sprintf("/config/namespaces/%s/api_credentials/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading API Credential",
			fmt.Sprintf("Could not read API credential %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
		return
	}

	data.Name = types.StringValue(apiResp.Metadata.Name)
	data.Namespace = types.StringValue(apiResp.Metadata.Namespace)
	if apiResp.Metadata.UID != "" {
		data.ID = types.StringValue(apiResp.Metadata.UID)
	} else {
		data.ID = types.StringValue(fmt.Sprintf("%s/%s", apiResp.Metadata.Namespace, apiResp.Metadata.Name))
	}
	if apiResp.Metadata.Description != "" {
		data.Description = types.StringValue(apiResp.Metadata.Description)
	}
	if apiResp.Spec.CredentialType != "" {
		data.CredentialType = types.StringValue(apiResp.Spec.CredentialType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
