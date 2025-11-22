// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package tcp_loadbalancer

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TCPLoadBalancerResourceModel describes the resource data model.
type TCPLoadBalancerResourceModel struct {
	// User-configurable fields
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`

	// Domains and ports
	Domains   []types.String `tfsdk:"domains"`
	ListenPort types.Int64   `tfsdk:"listen_port"`

	// Origin pool reference
	OriginPools []OriginPoolRefModel `tfsdk:"origin_pools"`

	// Advertisement
	AdvertiseOnPublic         types.Bool `tfsdk:"advertise_on_public"`
	AdvertiseOnPublicDefaultVIP types.Bool `tfsdk:"advertise_on_public_default_vip"`

	// DNS info
	DNSVolterraManaged types.Bool `tfsdk:"dns_volterra_managed"`

	// Load balancing options
	HashPolicyChoiceLeastActive types.Bool `tfsdk:"hash_policy_choice_least_active"`
	HashPolicyChoiceRandom      types.Bool `tfsdk:"hash_policy_choice_random"`
	HashPolicyChoiceSourceIP    types.Bool `tfsdk:"hash_policy_choice_source_ip"`

	// Connection settings
	IdleTimeout types.Int64 `tfsdk:"idle_timeout"`

	// Computed fields
	ID types.String `tfsdk:"id"`
}

// OriginPoolRefModel represents a reference to an origin pool.
type OriginPoolRefModel struct {
	PoolName      types.String `tfsdk:"pool_name"`
	PoolNamespace types.String `tfsdk:"pool_namespace"`
	Weight        types.Int64  `tfsdk:"weight"`
	Priority      types.Int64  `tfsdk:"priority"`
}

// API structures for F5 XC TCP Load Balancer

// APITCPLoadBalancer represents the API structure.
type APITCPLoadBalancer struct {
	Metadata   APIMetadata            `json:"metadata"`
	Spec       APITCPLoadBalancerSpec `json:"spec"`
	SystemMeta APISystemMetadata      `json:"system_metadata,omitempty"`
}

// APIMetadata represents metadata for API requests.
type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}

// APISystemMetadata represents system metadata from API.
type APISystemMetadata struct {
	UID               string `json:"uid,omitempty"`
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
}

// APITCPLoadBalancerSpec represents the spec for TCP Load Balancer.
type APITCPLoadBalancerSpec struct {
	Domains        []string                `json:"domains,omitempty"`
	ListenPort     uint32                  `json:"listen_port,omitempty"`
	OriginPools    []APIOriginPoolWithWeight `json:"origin_pools_weights,omitempty"`
	AdvertiseOnPublic         *struct{}   `json:"advertise_on_public,omitempty"`
	AdvertiseOnPublicDefaultVIP *struct{} `json:"advertise_on_public_default_vip,omitempty"`
	DNSVolterraManaged        *struct{}   `json:"dns_volterra_managed,omitempty"`
	AutomaticPort             *struct{}   `json:"automatic_port,omitempty"`
	NoSNI                     *struct{}   `json:"no_sni,omitempty"`
	HashPolicyChoiceLeastActive *struct{} `json:"hash_policy_choice_least_active,omitempty"`
	HashPolicyChoiceRandom      *struct{} `json:"hash_policy_choice_random,omitempty"`
	HashPolicyChoiceSourceIPStickiness *struct{} `json:"hash_policy_choice_source_ip_stickiness,omitempty"`
	IdleTimeout               uint32      `json:"idle_timeout,omitempty"`
}

// APIOriginPoolWithWeight represents an origin pool reference with weight.
type APIOriginPoolWithWeight struct {
	Pool     *APIPoolRef `json:"pool,omitempty"`
	Weight   uint32      `json:"weight,omitempty"`
	Priority uint32      `json:"priority,omitempty"`
}

// APIPoolRef represents a pool reference.
type APIPoolRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	Tenant    string `json:"tenant,omitempty"`
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *TCPLoadBalancerResourceModel) ToAPIRequest() *APITCPLoadBalancer {
	spec := APITCPLoadBalancerSpec{}

	// Set domains
	if len(m.Domains) > 0 {
		domains := make([]string, len(m.Domains))
		for i, d := range m.Domains {
			domains[i] = d.ValueString()
		}
		spec.Domains = domains
	}

	// Set listen port
	if !m.ListenPort.IsNull() && !m.ListenPort.IsUnknown() {
		spec.ListenPort = uint32(m.ListenPort.ValueInt64())
	}

	// Set origin pools
	if len(m.OriginPools) > 0 {
		pools := make([]APIOriginPoolWithWeight, len(m.OriginPools))
		for i, p := range m.OriginPools {
			pools[i] = APIOriginPoolWithWeight{
				Pool: &APIPoolRef{
					Name:      p.PoolName.ValueString(),
					Namespace: p.PoolNamespace.ValueString(),
				},
				Weight:   uint32(p.Weight.ValueInt64()),
				Priority: uint32(p.Priority.ValueInt64()),
			}
		}
		spec.OriginPools = pools
	}

	// Set advertisement
	if m.AdvertiseOnPublic.ValueBool() {
		spec.AdvertiseOnPublic = &struct{}{}
	}
	if m.AdvertiseOnPublicDefaultVIP.ValueBool() {
		spec.AdvertiseOnPublicDefaultVIP = &struct{}{}
	}

	// Set DNS
	if m.DNSVolterraManaged.ValueBool() {
		spec.DNSVolterraManaged = &struct{}{}
	}

	// Set hash policy
	if m.HashPolicyChoiceLeastActive.ValueBool() {
		spec.HashPolicyChoiceLeastActive = &struct{}{}
	}
	if m.HashPolicyChoiceRandom.ValueBool() {
		spec.HashPolicyChoiceRandom = &struct{}{}
	}
	if m.HashPolicyChoiceSourceIP.ValueBool() {
		spec.HashPolicyChoiceSourceIPStickiness = &struct{}{}
	}

	// Set idle timeout
	if !m.IdleTimeout.IsNull() && !m.IdleTimeout.IsUnknown() {
		spec.IdleTimeout = uint32(m.IdleTimeout.ValueInt64())
	}

	return &APITCPLoadBalancer{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

// FromAPIResponse updates the model from an API response.
func (m *TCPLoadBalancerResourceModel) FromAPIResponse(resp *APITCPLoadBalancer) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	// Set domains
	if len(resp.Spec.Domains) > 0 {
		domains := make([]types.String, len(resp.Spec.Domains))
		for i, d := range resp.Spec.Domains {
			domains[i] = types.StringValue(d)
		}
		m.Domains = domains
	}

	// Set listen port
	if resp.Spec.ListenPort > 0 {
		m.ListenPort = types.Int64Value(int64(resp.Spec.ListenPort))
	}

	// Set origin pools
	if len(resp.Spec.OriginPools) > 0 {
		pools := make([]OriginPoolRefModel, len(resp.Spec.OriginPools))
		for i, p := range resp.Spec.OriginPools {
			pool := OriginPoolRefModel{
				Weight:   types.Int64Value(int64(p.Weight)),
				Priority: types.Int64Value(int64(p.Priority)),
			}
			if p.Pool != nil {
				pool.PoolName = types.StringValue(p.Pool.Name)
				pool.PoolNamespace = types.StringValue(p.Pool.Namespace)
			}
			pools[i] = pool
		}
		m.OriginPools = pools
	}

	// Set advertisement
	m.AdvertiseOnPublic = types.BoolValue(resp.Spec.AdvertiseOnPublic != nil)
	m.AdvertiseOnPublicDefaultVIP = types.BoolValue(resp.Spec.AdvertiseOnPublicDefaultVIP != nil)

	// Set DNS
	m.DNSVolterraManaged = types.BoolValue(resp.Spec.DNSVolterraManaged != nil)

	// Set hash policy
	m.HashPolicyChoiceLeastActive = types.BoolValue(resp.Spec.HashPolicyChoiceLeastActive != nil)
	m.HashPolicyChoiceRandom = types.BoolValue(resp.Spec.HashPolicyChoiceRandom != nil)
	m.HashPolicyChoiceSourceIP = types.BoolValue(resp.Spec.HashPolicyChoiceSourceIPStickiness != nil)

	// Set idle timeout
	if resp.Spec.IdleTimeout > 0 {
		m.IdleTimeout = types.Int64Value(int64(resp.Spec.IdleTimeout))
	}

	// Set ID
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
