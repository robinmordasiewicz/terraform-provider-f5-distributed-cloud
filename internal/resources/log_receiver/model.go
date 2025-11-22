// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package log_receiver

import "github.com/hashicorp/terraform-plugin-framework/types"

type LogReceiverResourceModel struct {
	Name           types.String         `tfsdk:"name"`
	Namespace      types.String         `tfsdk:"namespace"`
	Description    types.String         `tfsdk:"description"`
	ReceiverType   types.String         `tfsdk:"receiver_type"`
	HTTPReceiver   *HTTPReceiverModel   `tfsdk:"http_receiver"`
	SyslogReceiver *SyslogReceiverModel `tfsdk:"syslog_receiver"`
	ID             types.String         `tfsdk:"id"`
}

type HTTPReceiverModel struct {
	URL       types.String    `tfsdk:"url"`
	AuthToken types.String    `tfsdk:"auth_token"`
	BatchSize types.Int64     `tfsdk:"batch_size"`
	TLSConfig *TLSConfigModel `tfsdk:"tls_config"`
}

type SyslogReceiverModel struct {
	Address  types.String `tfsdk:"address"`
	Port     types.Int64  `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
	Format   types.String `tfsdk:"format"`
}

type TLSConfigModel struct {
	SkipVerify types.Bool   `tfsdk:"skip_verify"`
	CACert     types.String `tfsdk:"ca_cert"`
}

type APILogReceiver struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APILogReceiverSpec `json:"spec"`
	SystemMeta APISystemMetadata  `json:"system_metadata,omitempty"`
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

type APILogReceiverSpec struct {
	ReceiverType   string             `json:"receiver_type,omitempty"`
	HTTPReceiver   *APIHTTPReceiver   `json:"http_receiver,omitempty"`
	SyslogReceiver *APISyslogReceiver `json:"syslog_receiver,omitempty"`
}

type APIHTTPReceiver struct {
	URL       string        `json:"url,omitempty"`
	AuthToken string        `json:"auth_token,omitempty"`
	BatchSize int64         `json:"batch_size,omitempty"`
	TLSConfig *APITLSConfig `json:"tls_config,omitempty"`
}

type APISyslogReceiver struct {
	Address  string `json:"address,omitempty"`
	Port     int64  `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Format   string `json:"format,omitempty"`
}

type APITLSConfig struct {
	SkipVerify bool   `json:"skip_verify,omitempty"`
	CACert     string `json:"ca_cert,omitempty"`
}

func (m *LogReceiverResourceModel) ToAPIRequest() *APILogReceiver {
	spec := APILogReceiverSpec{}
	if !m.ReceiverType.IsNull() {
		spec.ReceiverType = m.ReceiverType.ValueString()
	}
	if m.HTTPReceiver != nil {
		spec.HTTPReceiver = &APIHTTPReceiver{
			URL:       m.HTTPReceiver.URL.ValueString(),
			AuthToken: m.HTTPReceiver.AuthToken.ValueString(),
			BatchSize: m.HTTPReceiver.BatchSize.ValueInt64(),
		}
		if m.HTTPReceiver.TLSConfig != nil {
			spec.HTTPReceiver.TLSConfig = &APITLSConfig{
				SkipVerify: m.HTTPReceiver.TLSConfig.SkipVerify.ValueBool(),
				CACert:     m.HTTPReceiver.TLSConfig.CACert.ValueString(),
			}
		}
	}
	if m.SyslogReceiver != nil {
		spec.SyslogReceiver = &APISyslogReceiver{
			Address:  m.SyslogReceiver.Address.ValueString(),
			Port:     m.SyslogReceiver.Port.ValueInt64(),
			Protocol: m.SyslogReceiver.Protocol.ValueString(),
			Format:   m.SyslogReceiver.Format.ValueString(),
		}
	}
	return &APILogReceiver{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *LogReceiverResourceModel) FromAPIResponse(resp *APILogReceiver) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.Spec.ReceiverType != "" {
		m.ReceiverType = types.StringValue(resp.Spec.ReceiverType)
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
