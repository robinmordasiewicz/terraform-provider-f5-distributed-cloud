// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package global_log_receiver

import "github.com/hashicorp/terraform-plugin-framework/types"

type GlobalLogReceiverResourceModel struct {
	Name           types.String `tfsdk:"name"`
	Namespace      types.String `tfsdk:"namespace"`
	Description    types.String `tfsdk:"description"`
	ReceiverType   types.String `tfsdk:"receiver_type"`
	HTTPEndpoint   types.String `tfsdk:"http_endpoint"`
	SyslogHost     types.String `tfsdk:"syslog_host"`
	SyslogPort     types.Int64  `tfsdk:"syslog_port"`
	SplunkEndpoint types.String `tfsdk:"splunk_endpoint"`
	SplunkToken    types.String `tfsdk:"splunk_token"`
	ID             types.String `tfsdk:"id"`
}

type APIGlobalLogReceiver struct {
	Metadata   APIMetadata               `json:"metadata"`
	Spec       APIGlobalLogReceiverSpec  `json:"spec"`
	SystemMeta APISystemMetadata         `json:"system_metadata,omitempty"`
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

type APIGlobalLogReceiverSpec struct {
	ReceiverType   string `json:"receiver_type,omitempty"`
	HTTPEndpoint   string `json:"http_endpoint,omitempty"`
	SyslogHost     string `json:"syslog_host,omitempty"`
	SyslogPort     int64  `json:"syslog_port,omitempty"`
	SplunkEndpoint string `json:"splunk_endpoint,omitempty"`
	SplunkToken    string `json:"splunk_token,omitempty"`
}

func (m *GlobalLogReceiverResourceModel) ToAPIRequest() *APIGlobalLogReceiver {
	return &APIGlobalLogReceiver{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIGlobalLogReceiverSpec{
			ReceiverType:   m.ReceiverType.ValueString(),
			HTTPEndpoint:   m.HTTPEndpoint.ValueString(),
			SyslogHost:     m.SyslogHost.ValueString(),
			SyslogPort:     m.SyslogPort.ValueInt64(),
			SplunkEndpoint: m.SplunkEndpoint.ValueString(),
			SplunkToken:    m.SplunkToken.ValueString(),
		},
	}
}

func (m *GlobalLogReceiverResourceModel) FromAPIResponse(resp *APIGlobalLogReceiver) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.ReceiverType = types.StringValue(resp.Spec.ReceiverType)
	m.HTTPEndpoint = types.StringValue(resp.Spec.HTTPEndpoint)
	m.SyslogHost = types.StringValue(resp.Spec.SyslogHost)
	m.SyslogPort = types.Int64Value(resp.Spec.SyslogPort)
	m.SplunkEndpoint = types.StringValue(resp.Spec.SplunkEndpoint)
	m.SplunkToken = types.StringValue(resp.Spec.SplunkToken)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
