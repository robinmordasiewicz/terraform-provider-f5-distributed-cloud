// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package alert_receiver

import "github.com/hashicorp/terraform-plugin-framework/types"

type AlertReceiverResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Email       types.String `tfsdk:"email"`
	SlackURL    types.String `tfsdk:"slack_url"`
	PagerDuty   types.String `tfsdk:"pagerduty_key"`
	ID          types.String `tfsdk:"id"`
}

type APIAlertReceiver struct {
	Metadata   APIMetadata           `json:"metadata"`
	Spec       APIAlertReceiverSpec  `json:"spec"`
	SystemMeta APISystemMetadata     `json:"system_metadata,omitempty"`
}

type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}

type APISystemMetadata struct {
	UID               string `json:"uid,omitempty"`
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
}

type APIAlertReceiverSpec struct {
	Email     *APIEmailConfig     `json:"email,omitempty"`
	Slack     *APISlackConfig     `json:"slack,omitempty"`
	PagerDuty *APIPagerDutyConfig `json:"pagerduty,omitempty"`
}

type APIEmailConfig struct {
	To string `json:"to,omitempty"`
}

type APISlackConfig struct {
	URL string `json:"url,omitempty"`
}

type APIPagerDutyConfig struct {
	RoutingKey string `json:"routing_key,omitempty"`
}

func (m *AlertReceiverResourceModel) ToAPIRequest() *APIAlertReceiver {
	spec := APIAlertReceiverSpec{}
	if !m.Email.IsNull() && m.Email.ValueString() != "" {
		spec.Email = &APIEmailConfig{To: m.Email.ValueString()}
	}
	if !m.SlackURL.IsNull() && m.SlackURL.ValueString() != "" {
		spec.Slack = &APISlackConfig{URL: m.SlackURL.ValueString()}
	}
	if !m.PagerDuty.IsNull() && m.PagerDuty.ValueString() != "" {
		spec.PagerDuty = &APIPagerDutyConfig{RoutingKey: m.PagerDuty.ValueString()}
	}
	return &APIAlertReceiver{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *AlertReceiverResourceModel) FromAPIResponse(resp *APIAlertReceiver) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
	if resp.Spec.Email != nil {
		m.Email = types.StringValue(resp.Spec.Email.To)
	}
	if resp.Spec.Slack != nil {
		m.SlackURL = types.StringValue(resp.Spec.Slack.URL)
	}
	if resp.Spec.PagerDuty != nil {
		m.PagerDuty = types.StringValue(resp.Spec.PagerDuty.RoutingKey)
	}
}
