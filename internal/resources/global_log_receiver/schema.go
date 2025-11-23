// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package global_log_receiver

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages F5 Distributed Cloud Global Log Receiver.",
		MarkdownDescription: `
The ` + "`f5distributedcloud_global_log_receiver`" + ` resource manages Global Log Receivers in F5 Distributed Cloud.

Global Log Receivers enable forwarding of logs from all namespaces to external logging systems.

## Example Usage

` + "```hcl" + `
resource "f5distributedcloud_global_log_receiver" "example" {
  name           = "global-log-receiver"
  namespace      = "system"
  receiver_type  = "HTTP"
  http_endpoint  = "https://logs.example.com/ingest"
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
				Description: "Type of log receiver (HTTP, SYSLOG, SPLUNK, DATADOG).",
				Validators:  []validator.String{stringvalidator.OneOf("HTTP", "SYSLOG", "SPLUNK", "DATADOG")},
			},
			"http_endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "HTTP endpoint URL for log forwarding.",
			},
			"syslog_host": schema.StringAttribute{
				Optional:    true,
				Description: "Syslog server hostname or IP.",
			},
			"syslog_port": schema.Int64Attribute{
				Optional:    true,
				Description: "Syslog server port.",
			},
			"splunk_endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "Splunk HEC endpoint URL.",
			},
			"splunk_token": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Splunk HEC token.",
			},
		},
	}
}
