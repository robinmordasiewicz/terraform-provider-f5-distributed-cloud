// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package certificate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/client"
)

var _ datasource.DataSource = &CertificateDataSource{}

func NewCertificateDataSource() datasource.DataSource {
	return &CertificateDataSource{}
}

type CertificateDataSource struct {
	client *client.Client
}

type CertificateDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	CertType    types.String `tfsdk:"cert_type"`
}

type APICertificate struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		CertificateURL string      `json:"certificate_url,omitempty"`
		PrivateKey     interface{} `json:"private_key,omitempty"`
	} `json:"spec"`
}

func (d *CertificateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate"
}

func (d *CertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about an existing F5 Distributed Cloud Certificate.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_certificate`" + ` data source retrieves information about an existing certificate.

## Example Usage

` + "```hcl" + `
data "f5distributedcloud_certificate" "example" {
  name      = "my-certificate"
  namespace = "my-namespace"
}

output "cert_type" {
  value = data.f5distributedcloud_certificate.example.cert_type
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the certificate.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the certificate to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the certificate is located.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the certificate.",
				Computed:    true,
			},
			"cert_type": schema.StringAttribute{
				Description: "Type of certificate (TLS, MTLS).",
				Computed:    true,
			},
		},
	}
}

func (d *CertificateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *CertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CertificateDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APICertificate
	path := fmt.Sprintf("/config/namespaces/%s/certificates/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Certificate",
			fmt.Sprintf("Could not read certificate %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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

	// Determine certificate type
	if apiResp.Spec.PrivateKey != nil {
		data.CertType = types.StringValue("TLS")
	} else {
		data.CertType = types.StringValue("CA")
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
