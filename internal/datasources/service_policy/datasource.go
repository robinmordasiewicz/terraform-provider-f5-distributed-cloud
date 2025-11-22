// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package service_policy

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &ServicePolicyDataSource{}

func NewServicePolicyDataSource() datasource.DataSource {
	return &ServicePolicyDataSource{}
}

type ServicePolicyDataSource struct {
	client *client.Client
}

type ServicePolicyDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Algo        types.String `tfsdk:"algo"`
}

type APIServicePolicy struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		Algo string `json:"algo,omitempty"`
	} `json:"spec"`
}

func (d *ServicePolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_policy"
}

func (d *ServicePolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Service Policy.",
		MarkdownDescription: `
The ` + "`f5xc_service_policy`" + ` data source retrieves information about an existing service policy.

## Example Usage

` + "```hcl" + `
data "f5xc_service_policy" "example" {
  name      = "my-service-policy"
  namespace = "my-namespace"
}

output "policy_algo" {
  value = data.f5xc_service_policy.example.algo
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the service policy.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the service policy to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the service policy is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the service policy.",
				Computed:    true,
			},
			"algo": schema.StringAttribute{
				Description: "Algorithm used by the service policy (FIRST_MATCH, DENY_OVERRIDES, ALLOW_OVERRIDES).",
				Computed:    true,
			},
		},
	}
}

func (d *ServicePolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ServicePolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ServicePolicyDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIServicePolicy
	path := fmt.Sprintf("/config/namespaces/%s/service_policys/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Service Policy",
			fmt.Sprintf("Could not read service policy %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	if apiResp.Spec.Algo != "" {
		data.Algo = types.StringValue(apiResp.Spec.Algo)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
