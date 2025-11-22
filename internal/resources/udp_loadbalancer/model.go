// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package udp_loadbalancer

import "github.com/hashicorp/terraform-plugin-framework/types"

type UDPLoadBalancerResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Port        types.Int64  `tfsdk:"port"`
	ID          types.String `tfsdk:"id"`
}

type APIUDPLoadBalancer struct {
	Metadata   APIMetadata            `json:"metadata"`
	Spec       APIUDPLoadBalancerSpec `json:"spec"`
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

type APIUDPLoadBalancerSpec struct {
	Port int64 `json:"port,omitempty"`
}

func (m *UDPLoadBalancerResourceModel) ToAPIRequest() *APIUDPLoadBalancer {
	spec := APIUDPLoadBalancerSpec{}
	if !m.Port.IsNull() {
		spec.Port = m.Port.ValueInt64()
	}
	return &APIUDPLoadBalancer{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *UDPLoadBalancerResourceModel) FromAPIResponse(resp *APIUDPLoadBalancer) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.Spec.Port != 0 {
		m.Port = types.Int64Value(resp.Spec.Port)
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
