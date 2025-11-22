// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package k8s_cluster

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Kubernetes Cluster.",
		MarkdownDescription: `
The ` + "`f5xc_k8s_cluster`" + ` resource manages Kubernetes Clusters in F5 Distributed Cloud.

Kubernetes Clusters provide container orchestration capabilities.

## Example Usage

` + "```hcl" + `
resource "f5xc_k8s_cluster" "example" {
  name        = "my-k8s-cluster"
  namespace   = "system"
  description = "Production Kubernetes cluster"
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
