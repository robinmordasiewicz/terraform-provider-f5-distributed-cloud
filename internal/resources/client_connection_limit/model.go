// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package client_connection_limit

import "github.com/hashicorp/terraform-plugin-framework/types"

type ClientConnectionLimitResourceModel struct {
	Name               types.String `tfsdk:"name"`
	Namespace          types.String `tfsdk:"namespace"`
	Description        types.String `tfsdk:"description"`
	MaxConnections     types.Int64  `tfsdk:"max_connections"`
	ConnectionTimeout  types.Int64  `tfsdk:"connection_timeout"`
	EnableCircuitBreak types.Bool   `tfsdk:"enable_circuit_break"`
	ID                 types.String `tfsdk:"id"`
}

type APIClientConnectionLimit struct {
	Metadata   APIMetadata                   `json:"metadata"`
	Spec       APIClientConnectionLimitSpec  `json:"spec"`
	SystemMeta APISystemMetadata             `json:"system_metadata,omitempty"`
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

type APIClientConnectionLimitSpec struct {
	MaxConnections     int64 `json:"max_connections,omitempty"`
	ConnectionTimeout  int64 `json:"connection_timeout,omitempty"`
	EnableCircuitBreak bool  `json:"enable_circuit_break,omitempty"`
}

func (m *ClientConnectionLimitResourceModel) ToAPIRequest() *APIClientConnectionLimit {
	return &APIClientConnectionLimit{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIClientConnectionLimitSpec{
			MaxConnections:     m.MaxConnections.ValueInt64(),
			ConnectionTimeout:  m.ConnectionTimeout.ValueInt64(),
			EnableCircuitBreak: m.EnableCircuitBreak.ValueBool(),
		},
	}
}

func (m *ClientConnectionLimitResourceModel) FromAPIResponse(resp *APIClientConnectionLimit) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.MaxConnections = types.Int64Value(resp.Spec.MaxConnections)
	m.ConnectionTimeout = types.Int64Value(resp.Spec.ConnectionTimeout)
	m.EnableCircuitBreak = types.BoolValue(resp.Spec.EnableCircuitBreak)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
