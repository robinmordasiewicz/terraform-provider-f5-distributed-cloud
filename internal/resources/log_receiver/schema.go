// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package log_receiver

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Log Receiver.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_log_receiver`" + ` resource manages Log Receivers in F5 Distributed Cloud.

Log Receivers configure destinations for forwarding logs from F5 XC to external systems.

## Example Usage

### HTTP Log Receiver

` + "```hcl" + `
resource "f5_distributed_cloud_log_receiver" "http" {
  name          = "http-logs"
  namespace     = "system"
  receiver_type = "HTTP"

  http_receiver {
    url        = "https://logs.example.com/ingest"
    auth_token = var.log_token
    batch_size = 100

    tls_config {
      skip_verify = false
    }
  }
}
` + "```" + `

### Syslog Receiver

` + "```hcl" + `
resource "f5_distributed_cloud_log_receiver" "syslog" {
  name          = "syslog-logs"
  namespace     = "system"
  receiver_type = "SYSLOG"

  syslog_receiver {
    address  = "syslog.example.com"
    port     = 514
    protocol = "UDP"
    format   = "RFC5424"
  }
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":          schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":        schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":   schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description": schema.StringAttribute{Optional: true},
			"receiver_type": schema.StringAttribute{
				Required:    true,
				Description: "Type of log receiver.",
				Validators:  []validator.String{stringvalidator.OneOf("HTTP", "SYSLOG", "SPLUNK", "DATADOG", "AZURE_BLOB")},
			},
		},
		Blocks: map[string]schema.Block{
			"http_receiver": schema.SingleNestedBlock{
				Description: "HTTP receiver configuration.",
				Attributes: map[string]schema.Attribute{
					"url":        schema.StringAttribute{Required: true, Description: "HTTP endpoint URL for log delivery."},
					"auth_token": schema.StringAttribute{Optional: true, Sensitive: true, Description: "Authentication token for the HTTP endpoint."},
					"batch_size": schema.Int64Attribute{Optional: true, Description: "Number of logs to batch before sending."},
				},
				Blocks: map[string]schema.Block{
					"tls_config": schema.SingleNestedBlock{
						Description: "TLS configuration for HTTPS connections.",
						Attributes: map[string]schema.Attribute{
							"skip_verify": schema.BoolAttribute{Optional: true, Description: "Skip TLS certificate verification."},
							"ca_cert":     schema.StringAttribute{Optional: true, Sensitive: true, Description: "CA certificate for TLS verification."},
						},
					},
				},
			},
			"syslog_receiver": schema.SingleNestedBlock{
				Description: "Syslog receiver configuration.",
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{Required: true, Description: "Syslog server address."},
					"port":    schema.Int64Attribute{Required: true, Description: "Syslog server port."},
					"protocol": schema.StringAttribute{
						Optional:    true,
						Description: "Transport protocol (TCP/UDP).",
						Validators:  []validator.String{stringvalidator.OneOf("TCP", "UDP")},
					},
					"format": schema.StringAttribute{
						Optional:    true,
						Description: "Syslog message format.",
						Validators:  []validator.String{stringvalidator.OneOf("RFC3164", "RFC5424")},
					},
				},
			},
		},
	}
}
