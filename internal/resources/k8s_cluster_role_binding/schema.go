// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package k8s_cluster_role_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Kubernetes Cluster Role Binding.",
		MarkdownDescription: `
The ` + "`f5xc_k8s_cluster_role_binding`" + ` resource manages Kubernetes Cluster Role Bindings in F5 Distributed Cloud.

Kubernetes Cluster Role Bindings bind cluster roles to users or service accounts.

## Example Usage

` + "```hcl" + `
resource "f5xc_k8s_cluster_role_binding" "example" {
  name        = "my-cluster-role-binding"
  namespace   = "system"
  description = "Binding admin role to service account"
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
