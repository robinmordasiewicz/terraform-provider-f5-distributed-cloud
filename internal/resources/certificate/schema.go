// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// Schema returns the schema for the Certificate resource.
func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Certificate.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_certificate`" + ` resource manages TLS Certificates in F5 Distributed Cloud.

Certificates are used for TLS termination on load balancers and for securing
connections to origin servers.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_certificate" "example" {
  name        = "my-certificate"
  namespace   = "my-namespace"
  description = "My TLS certificate"

  certificate_url = "string:///LS0tLS1CRUdJTi..."

  private_key {
    blindfold_secret_url = "string:///..."
  }
}
` + "```" + `

## Import

Certificates can be imported using namespace/name:

` + "```shell" + `
terraform import f5distributedcloud_certificate.example my-namespace/my-certificate
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the Certificate.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Certificate.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the Certificate is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the Certificate.",
				Optional:    true,
			},
			"certificate_url": schema.StringAttribute{
				Description: "URL to the certificate in PEM format (base64 encoded with string:/// prefix).",
				Required:    true,
				Sensitive:   true,
			},
			"expiration": schema.StringAttribute{
				Description: "Certificate expiration date.",
				Computed:    true,
			},
			"subject": schema.StringAttribute{
				Description: "Certificate subject.",
				Computed:    true,
			},
			"issuer": schema.StringAttribute{
				Description: "Certificate issuer.",
				Computed:    true,
			},
		},
		Blocks: map[string]schema.Block{
			"private_key": schema.SingleNestedBlock{
				Description: "Private key configuration.",
				Attributes: map[string]schema.Attribute{
					"clear_secret_url": schema.StringAttribute{
						Description: "URL to clear text secret (not recommended for production).",
						Optional:    true,
						Sensitive:   true,
					},
					"blindfold_secret_url": schema.StringAttribute{
						Description: "URL to blindfolded secret.",
						Optional:    true,
						Sensitive:   true,
					},
					"vault_secret_url": schema.StringAttribute{
						Description: "URL to Vault secret.",
						Optional:    true,
						Sensitive:   true,
					},
					"wingman_secret_url": schema.StringAttribute{
						Description: "URL to Wingman secret.",
						Optional:    true,
						Sensitive:   true,
					},
				},
			},
		},
	}
}
