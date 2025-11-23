// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package alert_receiver

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func Schema() schema.Schema {
	return schema.Schema{
		Description: "Manages an F5 Distributed Cloud Alert Receiver.",
		MarkdownDescription: `
The ` + "`f5_distributed_cloud_alert_receiver`" + ` resource manages Alert Receivers in F5 Distributed Cloud.

Alert Receivers define notification destinations for alerts (email, Slack, PagerDuty).

## Example Usage

` + "```hcl" + `
resource "f5_distributed_cloud_alert_receiver" "slack" {
  name        = "slack-alerts"
  namespace   = "shared"
  description = "Slack channel for alerts"
  slack_url   = "https://hooks.slack.com/services/xxx/yyy/zzz"
}

resource "f5_distributed_cloud_alert_receiver" "email" {
  name        = "email-alerts"
  namespace   = "shared"
  description = "Email alerts"
  email       = "alerts@example.com"
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"id":            schema.StringAttribute{Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
			"name":          schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"namespace":     schema.StringAttribute{Required: true, PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()}},
			"description":   schema.StringAttribute{Optional: true},
			"email":         schema.StringAttribute{Optional: true, Description: "Email address to send alerts to."},
			"slack_url":     schema.StringAttribute{Optional: true, Sensitive: true, Description: "Slack webhook URL for notifications."},
			"pagerduty_key": schema.StringAttribute{Optional: true, Sensitive: true, Description: "PagerDuty routing key for notifications."},
		},
	}
}
