// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Tunnel.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_tunnel`" + ` resource manages Tunnels in F5 Distributed Cloud.

Tunnels provide secure connectivity between sites.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_tunnel" "example" {
  name            = "site-tunnel"
  namespace       = "my-namespace"
  tunnel_type     = "IPSEC"
  remote_ip       = "203.0.113.1"
  local_ip        = "192.0.2.1"
  encryption_type = "AES256"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"tunnel_type": schema.StringAttribute{
				Required:    true,
				Description: "Type of tunnel (IPSEC, GRE, VXLAN).",
				Validators:  []validator.String{stringvalidator.OneOf("IPSEC", "GRE", "VXLAN")},
			},
			"remote_ip": schema.StringAttribute{
				Required:    true,
				Description: "Remote endpoint IP address.",
			},
			"local_ip": schema.StringAttribute{
				Optional:    true,
				Description: "Local endpoint IP address.",
			},
			"encryption_type": schema.StringAttribute{
				Optional:    true,
				Description: "Encryption type for the tunnel (AES128, AES256).",
				Validators:  []validator.String{stringvalidator.OneOf("AES128", "AES256")},
			},
		},
	}
}
