// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_lb_pool

import "github.com/hashicorp/terraform-plugin-framework/types"

type DNSLBPoolResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	LoadBalancingMode types.String `tfsdk:"load_balancing_mode"`
	Members     []PoolMemberModel `tfsdk:"members"`
	ID          types.String `tfsdk:"id"`
}

type PoolMemberModel struct {
	Address  types.String `tfsdk:"address"`
	Port     types.Int64  `tfsdk:"port"`
	Priority types.Int64  `tfsdk:"priority"`
	Ratio    types.Int64  `tfsdk:"ratio"`
	Enabled  types.Bool   `tfsdk:"enabled"`
}

type APIDNSLBPool struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIDNSLBPoolSpec  `json:"spec"`
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

type APIDNSLBPoolSpec struct {
	LoadBalancingMode string          `json:"load_balancing_mode,omitempty"`
	Members           []APIPoolMember `json:"members,omitempty"`
}

type APIPoolMember struct {
	Address  string `json:"address,omitempty"`
	Port     int64  `json:"port,omitempty"`
	Priority int64  `json:"priority,omitempty"`
	Ratio    int64  `json:"ratio,omitempty"`
	Enabled  bool   `json:"enabled,omitempty"`
}

func (m *DNSLBPoolResourceModel) ToAPIRequest() *APIDNSLBPool {
	spec := APIDNSLBPoolSpec{}
	if !m.LoadBalancingMode.IsNull() {
		spec.LoadBalancingMode = m.LoadBalancingMode.ValueString()
	}
	if len(m.Members) > 0 {
		spec.Members = make([]APIPoolMember, len(m.Members))
		for i, member := range m.Members {
			spec.Members[i] = APIPoolMember{
				Address:  member.Address.ValueString(),
				Port:     member.Port.ValueInt64(),
				Priority: member.Priority.ValueInt64(),
				Ratio:    member.Ratio.ValueInt64(),
				Enabled:  member.Enabled.ValueBool(),
			}
		}
	}
	return &APIDNSLBPool{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *DNSLBPoolResourceModel) FromAPIResponse(resp *APIDNSLBPool) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.Spec.LoadBalancingMode != "" {
		m.LoadBalancingMode = types.StringValue(resp.Spec.LoadBalancingMode)
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
