// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package k8s_cluster_role

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Kubernetes Cluster Role.",
		MarkdownDescription: `
The ` + "`f5xc_k8s_cluster_role`" + ` resource manages Kubernetes Cluster Roles in F5 Distributed Cloud.

Kubernetes Cluster Roles define cluster-wide permissions for RBAC.

## Example Usage

` + "```hcl" + `
resource "f5xc_k8s_cluster_role" "example" {
  name        = "my-cluster-role"
  namespace   = "system"
  description = "Cluster role for admin access"
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
