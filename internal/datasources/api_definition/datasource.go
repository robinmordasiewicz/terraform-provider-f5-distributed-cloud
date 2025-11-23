// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package api_definition

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &APIDefinitionDataSource{}

func NewAPIDefinitionDataSource() datasource.DataSource {
	return &APIDefinitionDataSource{}
}

type APIDefinitionDataSource struct {
	client *client.Client
}

type APIDefinitionDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	SwaggerURL  types.String `tfsdk:"swagger_url"`
}

type APIAPIDefinition struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		SwaggerSpecs []struct {
			SwaggerURL string `json:"swagger_url,omitempty"`
		} `json:"swagger_specs,omitempty"`
	} `json:"spec"`
}

func (d *APIDefinitionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_definition"
}

func (d *APIDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud API Definition.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_api_definition`" + ` data source retrieves information about an existing API definition.

## Example Usage

` + "```hcl" + `
data "f5_distributed_cloud_api_definition" "example" {
  name      = "my-api-definition"
  namespace = "my-namespace"
}

output "swagger_url" {
  value = data.f5_distributed_cloud_api_definition.example.swagger_url
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the API definition.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the API definition to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the API definition is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the API definition.",
				Computed:    true,
			},
			"swagger_url": schema.StringAttribute{
				Description: "Swagger/OpenAPI specification URL.",
				Computed:    true,
			},
		},
	}
}

func (d *APIDefinitionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *APIDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data APIDefinitionDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIAPIDefinition
	path := fmt.Sprintf("/config/namespaces/%s/api_definitions/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading API Definition",
			fmt.Sprintf("Could not read API definition %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	if len(apiResp.Spec.SwaggerSpecs) > 0 && apiResp.Spec.SwaggerSpecs[0].SwaggerURL != "" {
		data.SwaggerURL = types.StringValue(apiResp.Spec.SwaggerSpecs[0].SwaggerURL)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
