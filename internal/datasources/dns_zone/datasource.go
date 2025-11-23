// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_zone

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &DNSZoneDataSource{}

func NewDNSZoneDataSource() datasource.DataSource {
	return &DNSZoneDataSource{}
}

type DNSZoneDataSource struct {
	client *client.Client
}

type DNSZoneDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Domain      types.String `tfsdk:"domain"`
}

type APIDNSZone struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		Domain string `json:"domain,omitempty"`
	} `json:"spec"`
}

func (d *DNSZoneDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_zone"
}

func (d *DNSZoneDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud DNS Zone.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_dns_zone`" + ` data source retrieves information about an existing DNS zone.

## Example Usage

` + "```hcl" + `
data "f5_distributed_cloud_dns_zone" "example" {
  name      = "my-dns-zone"
  namespace = "my-namespace"
}

output "domain" {
  value = data.f5_distributed_cloud_dns_zone.example.domain
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the DNS zone.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the DNS zone to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the DNS zone is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the DNS zone.",
				Computed:    true,
			},
			"domain": schema.StringAttribute{
				Description: "Domain name of the DNS zone.",
				Computed:    true,
			},
		},
	}
}

func (d *DNSZoneDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *DNSZoneDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DNSZoneDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIDNSZone
	path := fmt.Sprintf("/config/namespaces/%s/dns_zones/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading DNS Zone",
			fmt.Sprintf("Could not read DNS zone %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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
	if apiResp.Spec.Domain != "" {
		data.Domain = types.StringValue(apiResp.Spec.Domain)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
