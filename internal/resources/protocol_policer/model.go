// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package protocol_policer

import "github.com/hashicorp/terraform-plugin-framework/types"

type ProtocolPolicerResourceModel struct {
	Name            types.String `tfsdk:"name"`
	Namespace       types.String `tfsdk:"namespace"`
	Description     types.String `tfsdk:"description"`
	Protocol        types.String `tfsdk:"protocol"`
	RequestsPerUnit types.Int64  `tfsdk:"requests_per_unit"`
	Unit            types.String `tfsdk:"unit"`
	BurstSize       types.Int64  `tfsdk:"burst_size"`
	ID              types.String `tfsdk:"id"`
}

type APIProtocolPolicer struct {
	Metadata   APIMetadata            `json:"metadata"`
	Spec       APIProtocolPolicerSpec `json:"spec"`
	SystemMeta APISystemMetadata      `json:"system_metadata,omitempty"`
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

type APIProtocolPolicerSpec struct {
	Protocol        string `json:"protocol,omitempty"`
	RequestsPerUnit int64  `json:"requests_per_unit,omitempty"`
	Unit            string `json:"unit,omitempty"`
	BurstSize       int64  `json:"burst_size,omitempty"`
}

func (m *ProtocolPolicerResourceModel) ToAPIRequest() *APIProtocolPolicer {
	return &APIProtocolPolicer{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIProtocolPolicerSpec{
			Protocol:        m.Protocol.ValueString(),
			RequestsPerUnit: m.RequestsPerUnit.ValueInt64(),
			Unit:            m.Unit.ValueString(),
			BurstSize:       m.BurstSize.ValueInt64(),
		},
	}
}

func (m *ProtocolPolicerResourceModel) FromAPIResponse(resp *APIProtocolPolicer) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.Protocol = types.StringValue(resp.Spec.Protocol)
	m.RequestsPerUnit = types.Int64Value(resp.Spec.RequestsPerUnit)
	m.Unit = types.StringValue(resp.Spec.Unit)
	m.BurstSize = types.Int64Value(resp.Spec.BurstSize)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
