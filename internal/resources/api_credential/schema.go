// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package api_credential

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud API Credential.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_api_credential`" + ` resource manages API Credentials in F5 Distributed Cloud.

API Credentials are used for programmatic access to F5 XC APIs.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_api_credential" "example" {
  name            = "terraform-api-cred"
  namespace       = "system"
  credential_type = "API_CERTIFICATE"
  expiration_days = 365
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"credential_type": schema.StringAttribute{
				Required:   true,
				Validators: []validator.String{stringvalidator.OneOf("API_CERTIFICATE", "KUBE_CONFIG", "API_TOKEN")},
			},
			"expiration_days":       schema.Int64Attribute{Optional: true, Computed: true, Default: int64default.StaticInt64(365)},
			"virtual_k8s_name":      schema.StringAttribute{Optional: true, Description: "Name of the virtual K8s cluster (for KUBE_CONFIG type)"},
			"virtual_k8s_namespace": schema.StringAttribute{Optional: true, Description: "Namespace of the virtual K8s cluster"},
			"data":                  schema.StringAttribute{Computed: true, Sensitive: true, Description: "The credential data (certificate or token)"},
		},
	}
}
