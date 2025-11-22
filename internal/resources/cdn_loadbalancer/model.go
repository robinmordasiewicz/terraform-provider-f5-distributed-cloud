// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cdn_loadbalancer

import "github.com/hashicorp/terraform-plugin-framework/types"

type CDNLoadBalancerResourceModel struct {
	Name        types.String   `tfsdk:"name"`
	Namespace   types.String   `tfsdk:"namespace"`
	Description types.String   `tfsdk:"description"`
	Domains     []types.String `tfsdk:"domains"`
	ID          types.String   `tfsdk:"id"`
}

type APICDNLoadBalancer struct {
	Metadata   APIMetadata            `json:"metadata"`
	Spec       APICDNLoadBalancerSpec `json:"spec"`
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

type APICDNLoadBalancerSpec struct {
	Domains []string `json:"domains,omitempty"`
}

func (m *CDNLoadBalancerResourceModel) ToAPIRequest() *APICDNLoadBalancer {
	spec := APICDNLoadBalancerSpec{}
	if len(m.Domains) > 0 {
		domains := make([]string, len(m.Domains))
		for i, d := range m.Domains {
			domains[i] = d.ValueString()
		}
		spec.Domains = domains
	}
	return &APICDNLoadBalancer{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *CDNLoadBalancerResourceModel) FromAPIResponse(resp *APICDNLoadBalancer) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if len(resp.Spec.Domains) > 0 {
		domains := make([]types.String, len(resp.Spec.Domains))
		for i, d := range resp.Spec.Domains {
			domains[i] = types.StringValue(d)
		}
		m.Domains = domains
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
