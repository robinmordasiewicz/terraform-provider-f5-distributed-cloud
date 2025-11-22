// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package tunnel

import "github.com/hashicorp/terraform-plugin-framework/types"

type TunnelResourceModel struct {
	Name         types.String `tfsdk:"name"`
	Namespace    types.String `tfsdk:"namespace"`
	Description  types.String `tfsdk:"description"`
	TunnelType   types.String `tfsdk:"tunnel_type"`
	RemoteIP     types.String `tfsdk:"remote_ip"`
	LocalIP      types.String `tfsdk:"local_ip"`
	EncryptionType types.String `tfsdk:"encryption_type"`
	ID           types.String `tfsdk:"id"`
}

type APITunnel struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APITunnelSpec     `json:"spec"`
	SystemMeta APISystemMetadata `json:"system_metadata,omitempty"`
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

type APITunnelSpec struct {
	TunnelType     string `json:"tunnel_type,omitempty"`
	RemoteIP       string `json:"remote_ip,omitempty"`
	LocalIP        string `json:"local_ip,omitempty"`
	EncryptionType string `json:"encryption_type,omitempty"`
}

func (m *TunnelResourceModel) ToAPIRequest() *APITunnel {
	return &APITunnel{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APITunnelSpec{
			TunnelType:     m.TunnelType.ValueString(),
			RemoteIP:       m.RemoteIP.ValueString(),
			LocalIP:        m.LocalIP.ValueString(),
			EncryptionType: m.EncryptionType.ValueString(),
		},
	}
}

func (m *TunnelResourceModel) FromAPIResponse(resp *APITunnel) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.TunnelType = types.StringValue(resp.Spec.TunnelType)
	m.RemoteIP = types.StringValue(resp.Spec.RemoteIP)
	m.LocalIP = types.StringValue(resp.Spec.LocalIP)
	m.EncryptionType = types.StringValue(resp.Spec.EncryptionType)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
