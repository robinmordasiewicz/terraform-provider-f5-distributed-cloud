// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_load_balancer

import "github.com/hashicorp/terraform-plugin-framework/types"

type DNSLoadBalancerResourceModel struct {
	Name            types.String `tfsdk:"name"`
	Namespace       types.String `tfsdk:"namespace"`
	Description     types.String `tfsdk:"description"`
	Domain          types.String `tfsdk:"domain"`
	LoadBalanceType types.String `tfsdk:"load_balance_type"`
	ID              types.String `tfsdk:"id"`
}

type APIDNSLoadBalancer struct {
	Metadata   APIMetadata            `json:"metadata"`
	Spec       APIDNSLoadBalancerSpec `json:"spec"`
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

type APIDNSLoadBalancerSpec struct {
	Domain          string `json:"domain,omitempty"`
	LoadBalanceType string `json:"load_balance_type,omitempty"`
}

func (m *DNSLoadBalancerResourceModel) ToAPIRequest() *APIDNSLoadBalancer {
	return &APIDNSLoadBalancer{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     APIDNSLoadBalancerSpec{Domain: m.Domain.ValueString(), LoadBalanceType: m.LoadBalanceType.ValueString()},
	}
}

func (m *DNSLoadBalancerResourceModel) FromAPIResponse(resp *APIDNSLoadBalancer) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.Domain = types.StringValue(resp.Spec.Domain)
	m.LoadBalanceType = types.StringValue(resp.Spec.LoadBalanceType)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
