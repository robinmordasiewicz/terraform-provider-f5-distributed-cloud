// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cloud_credentials

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Cloud Credentials.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_cloud_credentials`" + ` resource manages Cloud Provider Credentials in F5 Distributed Cloud.

Cloud Credentials are used to authenticate with cloud providers (AWS, Azure, GCP) for site deployment and cloud resource management.

## Example Usage

### AWS Credentials

` + "```hcl" + `
resource "f5distributedcloud_cloud_credentials" "aws" {
  name           = "aws-creds"
  namespace      = "system"
  cloud_provider = "aws"

  aws_credentials {
    access_key_id     = var.aws_access_key
    secret_access_key = var.aws_secret_key
  }
}
` + "```" + `

### Azure Credentials

` + "```hcl" + `
resource "f5distributedcloud_cloud_credentials" "azure" {
  name           = "azure-creds"
  namespace      = "system"
  cloud_provider = "azure"

  azure_credentials {
    subscription_id = var.azure_subscription_id
    tenant_id       = var.azure_tenant_id
    client_id       = var.azure_client_id
    client_secret   = var.azure_client_secret
  }
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"cloud_provider": schema.StringAttribute{
				Required:   true,
				Validators: []validator.String{stringvalidator.OneOf("aws", "azure", "gcp")},
			},
		},
		Blocks: map[string]schema.Block{
			"aws_credentials": schema.SingleNestedBlock{
				Description: "AWS credentials configuration.",
				Attributes: map[string]schema.Attribute{
					"access_key_id":     schema.StringAttribute{Required: true, Sensitive: true},
					"secret_access_key": schema.StringAttribute{Required: true, Sensitive: true},
				},
			},
			"azure_credentials": schema.SingleNestedBlock{
				Description: "Azure credentials configuration.",
				Attributes: map[string]schema.Attribute{
					"subscription_id": schema.StringAttribute{Required: true},
					"tenant_id":       schema.StringAttribute{Required: true},
					"client_id":       schema.StringAttribute{Required: true},
					"client_secret":   schema.StringAttribute{Required: true, Sensitive: true},
				},
			},
			"gcp_credentials": schema.SingleNestedBlock{
				Description: "GCP credentials configuration.",
				Attributes: map[string]schema.Attribute{
					"project_id":       schema.StringAttribute{Required: true},
					"credentials_file": schema.StringAttribute{Required: true, Sensitive: true},
				},
			},
		},
	}
}
