// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package api_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud API Definition.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_api_definition`" + ` resource manages API Definitions in F5 Distributed Cloud.

API Definitions allow you to import and manage API schemas (OpenAPI/Swagger) for API security and discovery features.

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_api_definition" "example" {
  name        = "my-api"
  namespace   = "my-namespace"
  description = "My API definition"
  base_path   = "/api/v1"

  openapi_spec = <<-EOT
    openapi: "3.0.0"
    info:
      title: "My API"
      version: "1.0.0"
    paths:
      /users:
        get:
          summary: "List users"
  EOT
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"swagger_spec": schema.StringAttribute{
				Optional:    true,
				Description: "Swagger 2.0 specification in JSON or YAML format.",
			},
			"openapi_spec": schema.StringAttribute{
				Optional:    true,
				Description: "OpenAPI 3.0 specification in JSON or YAML format.",
			},
			"base_path": schema.StringAttribute{
				Optional:    true,
				Description: "Base path for the API endpoints.",
			},
		},
	}
}
