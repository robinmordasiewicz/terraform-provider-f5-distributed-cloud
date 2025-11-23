// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package service_mesh

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/client"
)

var _ datasource.DataSource = &ServiceMeshDataSource{}

func NewServiceMeshDataSource() datasource.DataSource {
	return &ServiceMeshDataSource{}
}

type ServiceMeshDataSource struct {
	client *client.Client
}

type ServiceMeshDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	MeshType    types.String `tfsdk:"mesh_type"`
}

type APIServiceMesh struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		MeshType string `json:"mesh_type,omitempty"`
	} `json:"spec"`
}

func (d *ServiceMeshDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_mesh"
}

func (d *ServiceMeshDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Service Mesh.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_service_mesh`" + ` data source retrieves information about an existing service mesh.

## Example Usage

` + "```hcl" + `
data "f5distributedcloud_service_mesh" "example" {
  name      = "my-service-mesh"
  namespace = "my-namespace"
}

output "mesh_type" {
  value = data.f5distributedcloud_service_mesh.example.mesh_type
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the service mesh.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the service mesh to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the service mesh is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the service mesh.",
				Computed:    true,
			},
			"mesh_type": schema.StringAttribute{
				Description: "Type of the service mesh.",
				Computed:    true,
			},
		},
	}
}

func (d *ServiceMeshDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ServiceMeshDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ServiceMeshDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIServiceMesh
	path := fmt.Sprintf("/config/namespaces/%s/service_meshs/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Service Mesh",
			fmt.Sprintf("Could not read service mesh %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	if apiResp.Spec.MeshType != "" {
		data.MeshType = types.StringValue(apiResp.Spec.MeshType)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
