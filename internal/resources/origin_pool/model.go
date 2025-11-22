// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package origin_pool

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// OriginPoolResourceModel describes the resource data model.
type OriginPoolResourceModel struct {
	// User-configurable fields
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`

	// Origin servers
	OriginServers []OriginServerModel `tfsdk:"origin_servers"`

	// Port and protocol
	Port            types.Int64  `tfsdk:"port"`
	EndpointProtocol types.String `tfsdk:"endpoint_protocol"`

	// Health check
	HealthCheckPort types.Int64 `tfsdk:"health_check_port"`

	// Load balancing
	LoadbalancerAlgorithm types.String `tfsdk:"loadbalancer_algorithm"`

	// Computed fields
	ID types.String `tfsdk:"id"`
}

// OriginServerModel represents an origin server configuration.
type OriginServerModel struct {
	PublicIP    *PublicIPModel     `tfsdk:"public_ip"`
	PrivateIP   *PrivateIPModel    `tfsdk:"private_ip"`
	PublicName  *PublicNameModel   `tfsdk:"public_name"`
	Labels      types.Map          `tfsdk:"labels"`
}

// PublicIPModel represents a public IP origin server.
type PublicIPModel struct {
	IP types.String `tfsdk:"ip"`
}

// PrivateIPModel represents a private IP origin server.
type PrivateIPModel struct {
	IP        types.String `tfsdk:"ip"`
	SiteName  types.String `tfsdk:"site_name"`
	InsideNetwork types.Bool `tfsdk:"inside_network"`
}

// PublicNameModel represents a DNS name origin server.
type PublicNameModel struct {
	DNSName types.String `tfsdk:"dns_name"`
}

// API structures for F5 XC Origin Pool

// APIOriginPool represents the API structure.
type APIOriginPool struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIOriginPoolSpec `json:"spec"`
	SystemMeta APISystemMetadata `json:"system_metadata,omitempty"`
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

// APIOriginPoolSpec represents the spec for Origin Pool.
type APIOriginPoolSpec struct {
	OriginServers         []APIOriginServer `json:"origin_servers,omitempty"`
	Port                  uint32            `json:"port,omitempty"`
	NoTLS                 *struct{}         `json:"no_tls,omitempty"`
	UseTLS                *APIUseTLS        `json:"use_tls,omitempty"`
	HealthCheckPort       uint32            `json:"healthcheck_port,omitempty"`
	LoadbalancerAlgorithm string            `json:"loadbalancer_algorithm,omitempty"`
}

// APIOriginServer represents an origin server in the API.
type APIOriginServer struct {
	PublicIP   *APIPublicIP   `json:"public_ip,omitempty"`
	PrivateIP  *APIPrivateIP  `json:"private_ip,omitempty"`
	PublicName *APIPublicName `json:"public_name,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
}

// APIPublicIP represents a public IP configuration.
type APIPublicIP struct {
	IP string `json:"ip"`
}

// APIPrivateIP represents a private IP configuration.
type APIPrivateIP struct {
	IP            string        `json:"ip"`
	SiteLocator   *APISiteLocator `json:"site_locator,omitempty"`
	InsideNetwork *struct{}     `json:"inside_network,omitempty"`
	OutsideNetwork *struct{}    `json:"outside_network,omitempty"`
}

// APISiteLocator represents a site locator.
type APISiteLocator struct {
	Site *APISiteRef `json:"site,omitempty"`
}

// APISiteRef represents a site reference.
type APISiteRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	Tenant    string `json:"tenant,omitempty"`
}

// APIPublicName represents a DNS name configuration.
type APIPublicName struct {
	DNSName string `json:"dns_name"`
}

// APIUseTLS represents TLS configuration.
type APIUseTLS struct {
	SkipServerVerification *struct{} `json:"skip_server_verification,omitempty"`
	UseHostHeaderAsSNI     *struct{} `json:"use_host_header_as_sni,omitempty"`
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *OriginPoolResourceModel) ToAPIRequest() *APIOriginPool {
	spec := APIOriginPoolSpec{
		Port: uint32(m.Port.ValueInt64()),
	}

	// Set endpoint protocol
	if m.EndpointProtocol.ValueString() == "https" {
		spec.UseTLS = &APIUseTLS{
			SkipServerVerification: &struct{}{},
		}
	} else {
		spec.NoTLS = &struct{}{}
	}

	// Set health check port
	if !m.HealthCheckPort.IsNull() && !m.HealthCheckPort.IsUnknown() {
		spec.HealthCheckPort = uint32(m.HealthCheckPort.ValueInt64())
	}

	// Set load balancer algorithm
	if !m.LoadbalancerAlgorithm.IsNull() && !m.LoadbalancerAlgorithm.IsUnknown() {
		spec.LoadbalancerAlgorithm = m.LoadbalancerAlgorithm.ValueString()
	}

	// Convert origin servers
	if len(m.OriginServers) > 0 {
		servers := make([]APIOriginServer, len(m.OriginServers))
		for i, s := range m.OriginServers {
			server := APIOriginServer{}
			if s.PublicIP != nil {
				server.PublicIP = &APIPublicIP{
					IP: s.PublicIP.IP.ValueString(),
				}
			}
			if s.PrivateIP != nil {
				server.PrivateIP = &APIPrivateIP{
					IP: s.PrivateIP.IP.ValueString(),
					SiteLocator: &APISiteLocator{
						Site: &APISiteRef{
							Name: s.PrivateIP.SiteName.ValueString(),
						},
					},
				}
				if s.PrivateIP.InsideNetwork.ValueBool() {
					server.PrivateIP.InsideNetwork = &struct{}{}
				} else {
					server.PrivateIP.OutsideNetwork = &struct{}{}
				}
			}
			if s.PublicName != nil {
				server.PublicName = &APIPublicName{
					DNSName: s.PublicName.DNSName.ValueString(),
				}
			}
			servers[i] = server
		}
		spec.OriginServers = servers
	}

	return &APIOriginPool{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

// FromAPIResponse updates the model from an API response.
func (m *OriginPoolResourceModel) FromAPIResponse(resp *APIOriginPool) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.Port = types.Int64Value(int64(resp.Spec.Port))

	// Set endpoint protocol
	if resp.Spec.UseTLS != nil {
		m.EndpointProtocol = types.StringValue("https")
	} else {
		m.EndpointProtocol = types.StringValue("http")
	}

	// Set health check port
	if resp.Spec.HealthCheckPort > 0 {
		m.HealthCheckPort = types.Int64Value(int64(resp.Spec.HealthCheckPort))
	}

	// Set load balancer algorithm
	if resp.Spec.LoadbalancerAlgorithm != "" {
		m.LoadbalancerAlgorithm = types.StringValue(resp.Spec.LoadbalancerAlgorithm)
	}

	// Convert origin servers
	if len(resp.Spec.OriginServers) > 0 {
		servers := make([]OriginServerModel, len(resp.Spec.OriginServers))
		for i, s := range resp.Spec.OriginServers {
			server := OriginServerModel{}
			if s.PublicIP != nil {
				server.PublicIP = &PublicIPModel{
					IP: types.StringValue(s.PublicIP.IP),
				}
			}
			if s.PrivateIP != nil {
				siteName := ""
				if s.PrivateIP.SiteLocator != nil && s.PrivateIP.SiteLocator.Site != nil {
					siteName = s.PrivateIP.SiteLocator.Site.Name
				}
				server.PrivateIP = &PrivateIPModel{
					IP:            types.StringValue(s.PrivateIP.IP),
					SiteName:      types.StringValue(siteName),
					InsideNetwork: types.BoolValue(s.PrivateIP.InsideNetwork != nil),
				}
			}
			if s.PublicName != nil {
				server.PublicName = &PublicNameModel{
					DNSName: types.StringValue(s.PublicName.DNSName),
				}
			}
			servers[i] = server
		}
		m.OriginServers = servers
	}

	// Set ID
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
