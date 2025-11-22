// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package network_firewall

import "github.com/hashicorp/terraform-plugin-framework/types"

type NetworkFirewallResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
}

type APINetworkFirewall struct {
	Metadata   APIMetadata            `json:"metadata"`
	Spec       APINetworkFirewallSpec `json:"spec"`
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

type APINetworkFirewallSpec struct {
	// Network firewall configuration fields
}

func (m *NetworkFirewallResourceModel) ToAPIRequest() *APINetworkFirewall {
	return &APINetworkFirewall{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     APINetworkFirewallSpec{},
	}
}

func (m *NetworkFirewallResourceModel) FromAPIResponse(resp *APINetworkFirewall) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
