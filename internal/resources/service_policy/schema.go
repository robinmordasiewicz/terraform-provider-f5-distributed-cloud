// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package service_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Schema returns the schema for the Service Policy resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Service Policy.",
		MarkdownDescription: `
The ` + "`f5xc_service_policy`" + ` resource manages Service Policies in F5 Distributed Cloud.

Service Policies define traffic control rules for service-to-service communication.
They can be used to implement security policies, traffic routing, and access control.

## Example Usage

` + "```hcl" + `
resource "f5xc_service_policy" "example" {
  name        = "my-service-policy"
  namespace   = "my-namespace"
  description = "Example service policy"

  default_action = "deny"

  rules {
    name   = "allow-internal"
    action = "allow"
  }
}
` + "```" + `

## Import

Service Policies can be imported using namespace/name:

` + "```shell" + `
terraform import f5xc_service_policy.example my-namespace/my-service-policy
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Service Policy.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Service Policy.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Service Policy is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Service Policy.",
				Optional:    true,
			},
			"default_action": schema.StringAttribute{
				Description: "Default action when no rules match: 'allow' or 'deny'.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("deny"),
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "deny"),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"rules": schema.ListNestedBlock{
				Description: "List of policy rules.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Name of the rule.",
							Required:    true,
						},
						"action": schema.StringAttribute{
							Description: "Action for the rule: 'allow' or 'deny'.",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "deny"),
							},
						},
					},
					Blocks: map[string]schema.Block{
						"match_sources": schema.ListNestedBlock{
							Description: "Source matching criteria.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										Description: "IP prefix to match (CIDR notation).",
										Optional:    true,
									},
									"prefix_list": schema.StringAttribute{
										Description: "Reference to a prefix list.",
										Optional:    true,
									},
									"asn": schema.Int64Attribute{
										Description: "ASN to match.",
										Optional:    true,
									},
									"namespace": schema.StringAttribute{
										Description: "Namespace to match.",
										Optional:    true,
									},
								},
							},
						},
						"match_destinations": schema.ListNestedBlock{
							Description: "Destination matching criteria.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										Description: "IP prefix to match (CIDR notation).",
										Optional:    true,
									},
									"port": schema.Int64Attribute{
										Description: "Port to match.",
										Optional:    true,
									},
									"protocol": schema.StringAttribute{
										Description: "Protocol to match: 'tcp', 'udp', 'icmp'.",
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
