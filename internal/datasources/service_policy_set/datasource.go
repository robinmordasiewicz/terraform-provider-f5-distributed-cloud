// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package service_policy_set

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &ServicePolicySetDataSource{}

func NewServicePolicySetDataSource() datasource.DataSource {
	return &ServicePolicySetDataSource{}
}

type ServicePolicySetDataSource struct {
	client *client.Client
}

type ServicePolicySetDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

type APIServicePolicySet struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace,omitempty"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
}

func (d *ServicePolicySetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_policy_set"
}

func (d *ServicePolicySetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Service Policy Set.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the service policy set.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the service policy set to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the service policy set exists.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the service policy set.",
				Computed:    true,
			},
		},
	}
}

func (d *ServicePolicySetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ServicePolicySetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ServicePolicySetDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIServicePolicySet
	path := fmt.Sprintf("/config/namespaces/%s/service_policy_sets/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Service Policy Set",
			fmt.Sprintf("Could not read service policy set %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
		return
	}

	data.Name = types.StringValue(apiResp.Metadata.Name)
	data.Namespace = types.StringValue(apiResp.Metadata.Namespace)
	if apiResp.Metadata.UID != "" {
		data.ID = types.StringValue(apiResp.Metadata.UID)
	} else {
		data.ID = types.StringValue(apiResp.Metadata.Namespace + "/" + apiResp.Metadata.Name)
	}
	if apiResp.Metadata.Description != "" {
		data.Description = types.StringValue(apiResp.Metadata.Description)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
