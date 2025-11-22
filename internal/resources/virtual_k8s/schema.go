// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package virtual_k8s

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Virtual Kubernetes.",
		MarkdownDescription: `
The ` + "`f5xc_virtual_k8s`" + ` resource manages Virtual Kubernetes clusters in F5 Distributed Cloud.

Virtual Kubernetes provides managed Kubernetes namespaces across distributed sites.

## Example Usage

` + "```hcl" + `
resource "f5xc_virtual_k8s" "example" {
  name        = "my-vk8s"
  namespace   = "system"
  description = "Virtual Kubernetes for application workloads"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
		},
	}
}
