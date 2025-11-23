// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package discovery

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Discovery configuration.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_discovery`" + ` resource manages Discovery configurations in F5 Distributed Cloud.

Discovery enables automatic detection of services running in Kubernetes clusters or other environments connected to F5 XC.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_discovery" "k8s_discovery" {
  name           = "k8s-discovery"
  namespace      = "my-namespace"
  discovery_type = "K8S"
  cluster_name   = "my-k8s-cluster"

  publish_info {
    disable_publish = false
    dns_domain      = "services.example.com"
  }
}
` + "```" + `

## Virtual Site Discovery

` + "```hcl" + `
resource "f5distributedcloud_discovery" "vsite_discovery" {
  name             = "vsite-discovery"
  namespace        = "my-namespace"
  discovery_type   = "VIRTUAL_SITE"
  virtual_site_ref = "my-virtual-site"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"discovery_type": schema.StringAttribute{
				Required:    true,
				Description: "Type of discovery to perform.",
				Validators:  []validator.String{stringvalidator.OneOf("K8S", "CONSUL", "VIRTUAL_SITE")},
			},
			"cluster_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the Kubernetes cluster for K8S discovery.",
			},
			"kube_config": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Kubernetes configuration for cluster access.",
			},
			"virtual_site_ref": schema.StringAttribute{
				Optional:    true,
				Description: "Reference to virtual site for VIRTUAL_SITE discovery type.",
			},
		},
		Blocks: map[string]schema.Block{
			"publish_info": schema.SingleNestedBlock{
				Description: "Service publishing configuration.",
				Attributes: map[string]schema.Attribute{
					"disable_publish": schema.BoolAttribute{
						Optional:    true,
						Description: "Disable automatic service publishing.",
					},
					"dns_domain": schema.StringAttribute{
						Optional:    true,
						Description: "DNS domain for discovered services.",
					},
				},
			},
		},
	}
}
