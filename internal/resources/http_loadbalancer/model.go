// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package http_loadbalancer

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HTTPLoadBalancerResourceModel describes the resource data model.
type HTTPLoadBalancerResourceModel struct {
	// User-configurable fields
	Name        types.String   `tfsdk:"name"`
	Namespace   types.String   `tfsdk:"namespace"`
	Description types.String   `tfsdk:"description"`
	Domains     []types.String `tfsdk:"domains"`

	// Load balancer type - HTTP or HTTPS
	HTTPType    types.String   `tfsdk:"http_type"`

	// Port configuration
	AdvertisePort types.Int64  `tfsdk:"advertise_port"`

	// Origin pools
	DefaultOriginPools []OriginPoolModel `tfsdk:"default_origin_pools"`

	// WAF configuration
	DisableWAF  types.Bool   `tfsdk:"disable_waf"`

	// Computed fields
	ID types.String `tfsdk:"id"`
}

// OriginPoolModel represents an origin pool reference.
type OriginPoolModel struct {
	PoolName   types.String `tfsdk:"pool_name"`
	PoolNamespace types.String `tfsdk:"pool_namespace"`
	Weight     types.Int64  `tfsdk:"weight"`
	Priority   types.Int64  `tfsdk:"priority"`
}

// API structures for F5 XC HTTP Load Balancer

// APIHTTPLoadBalancer represents the API structure for HTTP Load Balancer.
type APIHTTPLoadBalancer struct {
	Metadata   APIMetadata          `json:"metadata"`
	Spec       APIHTTPLoadBalancerSpec `json:"spec"`
	SystemMeta APISystemMetadata    `json:"system_metadata,omitempty"`
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

// APIHTTPLoadBalancerSpec represents the spec for HTTP Load Balancer.
type APIHTTPLoadBalancerSpec struct {
	Domains            []string                 `json:"domains,omitempty"`
	HTTP               *APIHTTPConfig           `json:"http,omitempty"`
	HTTPS              *APIHTTPSConfig          `json:"https,omitempty"`
	AdvertiseCustom    *APIAdvertiseCustom      `json:"advertise_custom,omitempty"`
	DefaultOriginPools []APIOriginPoolRef       `json:"default_origin_servers,omitempty"`
	DisableWAF         *struct{}                `json:"disable_waf,omitempty"`
	AppFirewall        *APIAppFirewallRef       `json:"app_firewall,omitempty"`
}

// APIHTTPConfig represents HTTP configuration.
type APIHTTPConfig struct {
	Port uint32 `json:"port,omitempty"`
}

// APIHTTPSConfig represents HTTPS configuration.
type APIHTTPSConfig struct {
	HTTPRedirect bool   `json:"http_redirect,omitempty"`
	Port         uint32 `json:"port,omitempty"`
}

// APIAdvertiseCustom represents custom advertisement configuration.
type APIAdvertiseCustom struct {
	AdvertiseWhere []APIAdvertiseWhere `json:"advertise_where,omitempty"`
}

// APIAdvertiseWhere represents where to advertise.
type APIAdvertiseWhere struct {
	Site       *APISiteRef `json:"site,omitempty"`
	VirtualSite *APISiteRef `json:"virtual_site,omitempty"`
	Port       uint32      `json:"port,omitempty"`
}

// APISiteRef represents a site reference.
type APISiteRef struct {
	Site      *APIObjectRef `json:"site,omitempty"`
	Network   string        `json:"network,omitempty"`
}

// APIObjectRef represents an object reference.
type APIObjectRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Tenant    string `json:"tenant,omitempty"`
}

// APIOriginPoolRef represents an origin pool reference.
type APIOriginPoolRef struct {
	Pool     *APIObjectRef `json:"pool,omitempty"`
	Weight   uint32        `json:"weight,omitempty"`
	Priority uint32        `json:"priority,omitempty"`
}

// APIAppFirewallRef represents an app firewall reference.
type APIAppFirewallRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *HTTPLoadBalancerResourceModel) ToAPIRequest() *APIHTTPLoadBalancer {
	// Build domains slice
	domains := make([]string, len(m.Domains))
	for i, d := range m.Domains {
		domains[i] = d.ValueString()
	}

	spec := APIHTTPLoadBalancerSpec{
		Domains: domains,
	}

	// Set HTTP type configuration
	if m.HTTPType.ValueString() == "https" {
		spec.HTTPS = &APIHTTPSConfig{
			HTTPRedirect: true,
		}
	} else {
		port := uint32(80)
		if !m.AdvertisePort.IsNull() && !m.AdvertisePort.IsUnknown() {
			port = uint32(m.AdvertisePort.ValueInt64())
		}
		spec.HTTP = &APIHTTPConfig{
			Port: port,
		}
	}

	// Set origin pools
	if len(m.DefaultOriginPools) > 0 {
		pools := make([]APIOriginPoolRef, len(m.DefaultOriginPools))
		for i, p := range m.DefaultOriginPools {
			pools[i] = APIOriginPoolRef{
				Pool: &APIObjectRef{
					Name:      p.PoolName.ValueString(),
					Namespace: p.PoolNamespace.ValueString(),
				},
				Weight:   uint32(p.Weight.ValueInt64()),
				Priority: uint32(p.Priority.ValueInt64()),
			}
		}
		spec.DefaultOriginPools = pools
	}

	// Set WAF configuration
	if m.DisableWAF.ValueBool() {
		spec.DisableWAF = &struct{}{}
	}

	return &APIHTTPLoadBalancer{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

// FromAPIResponse updates the model from an API response.
func (m *HTTPLoadBalancerResourceModel) FromAPIResponse(resp *APIHTTPLoadBalancer) {
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

	// Determine HTTP type
	if resp.Spec.HTTPS != nil {
		m.HTTPType = types.StringValue("https")
	} else {
		m.HTTPType = types.StringValue("http")
		if resp.Spec.HTTP != nil {
			m.AdvertisePort = types.Int64Value(int64(resp.Spec.HTTP.Port))
		}
	}

	// Set origin pools
	if len(resp.Spec.DefaultOriginPools) > 0 {
		pools := make([]OriginPoolModel, len(resp.Spec.DefaultOriginPools))
		for i, p := range resp.Spec.DefaultOriginPools {
			pools[i] = OriginPoolModel{
				PoolName:      types.StringValue(p.Pool.Name),
				PoolNamespace: types.StringValue(p.Pool.Namespace),
				Weight:        types.Int64Value(int64(p.Weight)),
				Priority:      types.Int64Value(int64(p.Priority)),
			}
		}
		m.DefaultOriginPools = pools
	}

	// Set WAF status
	m.DisableWAF = types.BoolValue(resp.Spec.DisableWAF != nil)

	// Set ID
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
