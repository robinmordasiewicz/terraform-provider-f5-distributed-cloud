// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cloud_credentials

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ datasource.DataSource = &CloudCredentialsDataSource{}

func NewCloudCredentialsDataSource() datasource.DataSource {
	return &CloudCredentialsDataSource{}
}

type CloudCredentialsDataSource struct {
	client *client.Client
}

type CloudCredentialsDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	CloudType   types.String `tfsdk:"cloud_type"`
}

type APICloudCredentials struct {
	Metadata struct {
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		Description string `json:"description,omitempty"`
		UID         string `json:"uid,omitempty"`
	} `json:"metadata"`
	Spec struct {
		AWSSecretKey      interface{} `json:"aws_secret_key,omitempty"`
		AzureClientSecret interface{} `json:"azure_client_secret,omitempty"`
		GCPCredFile       interface{} `json:"gcp_cred_file,omitempty"`
	} `json:"spec"`
}

func (d *CloudCredentialsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_credentials"
}

func (d *CloudCredentialsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches information about existing F5 Distributed Cloud Cloud Credentials.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_cloud_credentials`" + ` data source retrieves information about existing cloud credentials.

## Example Usage

` + "```hcl" + `
data "f5_distributed_cloud_cloud_credentials" "example" {
  name      = "my-aws-creds"
  namespace = "system"
}

output "cloud_type" {
  value = data.f5_distributed_cloud_cloud_credentials.example.cloud_type
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier of the cloud credentials.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the cloud credentials to look up.",
				Required:    true,
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the cloud credentials are located (typically 'system').",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the cloud credentials.",
				Computed:    true,
			},
			"cloud_type": schema.StringAttribute{
				Description: "Type of cloud provider (aws, azure, gcp).",
				Computed:    true,
			},
		},
	}
}

func (d *CloudCredentialsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *CloudCredentialsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CloudCredentialsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APICloudCredentials
	path := fmt.Sprintf("/config/namespaces/%s/cloud_credentialss/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := d.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Reading Cloud Credentials",
			fmt.Sprintf("Could not read cloud credentials %s/%s: %s", data.Namespace.ValueString(), data.Name.ValueString(), err))
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

	// Determine cloud type based on which credential is present
	cloudType := "unknown"
	if apiResp.Spec.AWSSecretKey != nil {
		cloudType = "aws"
	} else if apiResp.Spec.AzureClientSecret != nil {
		cloudType = "azure"
	} else if apiResp.Spec.GCPCredFile != nil {
		cloudType = "gcp"
	}
	data.CloudType = types.StringValue(cloudType)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
