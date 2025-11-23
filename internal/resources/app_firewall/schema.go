// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package app_firewall

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Schema returns the schema for the App Firewall resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Application Firewall (WAF) policy.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_app_firewall`" + ` resource manages Application Firewall (WAF) policies in F5 Distributed Cloud.

Application Firewall policies define web application security settings including
detection modes, bot protection, and custom blocking pages.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_app_firewall" "example" {
  name        = "my-app-firewall"
  namespace   = "my-namespace"
  description = "Example Application Firewall"
  mode        = "blocking"

  detection_settings {
    signature_based_bot_protection = true
    xhr_check                      = true
    cookie_protection              = true
    violation_rating               = "medium"
  }

  bot_protection {
    good_bot_action       = "allow"
    malicious_bot_action  = "block"
    suspicious_bot_action = "challenge"
  }
}
` + "```" + `

## Import

App Firewalls can be imported using namespace/name:

` + "```shell" + `
terraform import f5_distributed_cloud_app_firewall.example my-namespace/my-app-firewall
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the App Firewall.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the App Firewall.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the App Firewall is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the App Firewall.",
				Optional:    true,
			},
			"mode": schema.StringAttribute{
				Description: "Firewall mode: 'monitoring' or 'blocking'.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("monitoring"),
				Validators: []validator.String{
					stringvalidator.OneOf("monitoring", "blocking"),
				},
			},
			"allowed_response_codes": schema.ListAttribute{
				Description: "List of HTTP response codes that are allowed.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"default_anonymization": schema.BoolAttribute{
				Description: "Enable default anonymization of sensitive data.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"use_default_blocking_page": schema.BoolAttribute{
				Description: "Use the default blocking page for blocked requests.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"labels": schema.MapAttribute{
				Description: "Labels to apply to the App Firewall.",
				Optional:    true,
				ElementType: types.StringType,
			},
		},
		Blocks: map[string]schema.Block{
			"detection_settings": schema.SingleNestedBlock{
				Description: "Detection settings for the firewall.",
				Attributes: map[string]schema.Attribute{
					"signature_based_bot_protection": schema.BoolAttribute{
						Description: "Enable signature-based bot protection.",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(true),
					},
					"xhr_check": schema.BoolAttribute{
						Description: "Enable XHR request checking.",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
					},
					"cookie_protection": schema.BoolAttribute{
						Description: "Enable cookie protection.",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
					},
					"enable_suppression": schema.BoolAttribute{
						Description: "Enable violation suppression.",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
					},
					"violation_rating": schema.StringAttribute{
						Description: "Violation rating threshold: 'low', 'medium', or 'high'.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("medium"),
						Validators: []validator.String{
							stringvalidator.OneOf("low", "medium", "high"),
						},
					},
				},
			},
			"bot_protection": schema.SingleNestedBlock{
				Description: "Bot protection settings.",
				Attributes: map[string]schema.Attribute{
					"good_bot_action": schema.StringAttribute{
						Description: "Action for good bots: 'allow', 'block', or 'challenge'.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("allow"),
						Validators: []validator.String{
							stringvalidator.OneOf("allow", "block", "challenge"),
						},
					},
					"malicious_bot_action": schema.StringAttribute{
						Description: "Action for malicious bots: 'allow', 'block', or 'challenge'.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("block"),
						Validators: []validator.String{
							stringvalidator.OneOf("allow", "block", "challenge"),
						},
					},
					"suspicious_bot_action": schema.StringAttribute{
						Description: "Action for suspicious bots: 'allow', 'block', or 'challenge'.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("challenge"),
						Validators: []validator.String{
							stringvalidator.OneOf("allow", "block", "challenge"),
						},
					},
				},
			},
			"custom_blocking_page": schema.SingleNestedBlock{
				Description: "Custom blocking page configuration.",
				Attributes: map[string]schema.Attribute{
					"blocking_page_body": schema.StringAttribute{
						Description: "HTML body content for the custom blocking page.",
						Required:    true,
					},
					"response_code": schema.StringAttribute{
						Description: "HTTP response code for the blocking page.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("403"),
					},
				},
			},
		},
	}
}
