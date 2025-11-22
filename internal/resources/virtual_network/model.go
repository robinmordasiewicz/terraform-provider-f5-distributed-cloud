// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package virtual_network

import "github.com/hashicorp/terraform-plugin-framework/types"

type VirtualNetworkResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
}

type APIVirtualNetwork struct {
	Metadata   APIMetadata           `json:"metadata"`
	Spec       APIVirtualNetworkSpec `json:"spec"`
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

type APIVirtualNetworkSpec struct{}

func (m *VirtualNetworkResourceModel) ToAPIRequest() *APIVirtualNetwork {
	return &APIVirtualNetwork{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     APIVirtualNetworkSpec{},
	}
}

func (m *VirtualNetworkResourceModel) FromAPIResponse(resp *APIVirtualNetwork) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
