// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package healthcheck

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// HealthcheckResourceModel describes the resource data model.
type HealthcheckResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`

	// Health check type (only one should be set)
	HTTPHealthCheck *HTTPHealthCheckModel `tfsdk:"http_health_check"`
	TCPHealthCheck  *TCPHealthCheckModel  `tfsdk:"tcp_health_check"`

	// Common settings
	Timeout           types.Int64 `tfsdk:"timeout"`
	Interval          types.Int64 `tfsdk:"interval"`
	UnhealthyThreshold types.Int64 `tfsdk:"unhealthy_threshold"`
	HealthyThreshold   types.Int64 `tfsdk:"healthy_threshold"`

	// Computed fields
	ID types.String `tfsdk:"id"`
}

// HTTPHealthCheckModel represents HTTP health check settings.
type HTTPHealthCheckModel struct {
	Path               types.String `tfsdk:"path"`
	UseHTTPS           types.Bool   `tfsdk:"use_https"`
	ExpectedStatusCodes []types.String `tfsdk:"expected_status_codes"`
	Headers            types.Map    `tfsdk:"headers"`
}

// TCPHealthCheckModel represents TCP health check settings.
type TCPHealthCheckModel struct {
	// TCP health checks typically just check connectivity
}

// API structures

type APIHealthcheck struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APIHealthcheckSpec `json:"spec"`
	SystemMeta APISystemMetadata  `json:"system_metadata,omitempty"`
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

type APIHealthcheckSpec struct {
	HTTPHealthCheck    *APIHTTPHealthCheck `json:"http_health_check,omitempty"`
	TCPHealthCheck     *APITCPHealthCheck  `json:"tcp_health_check,omitempty"`
	Timeout            uint32              `json:"timeout,omitempty"`
	Interval           uint32              `json:"interval,omitempty"`
	UnhealthyThreshold uint32              `json:"unhealthy_threshold,omitempty"`
	HealthyThreshold   uint32              `json:"healthy_threshold,omitempty"`
}

type APIHTTPHealthCheck struct {
	Path                string            `json:"path,omitempty"`
	UseHTTPS            *struct{}         `json:"use_https,omitempty"`
	ExpectedStatusCodes []string          `json:"expected_status_codes,omitempty"`
	Headers             map[string]string `json:"headers,omitempty"`
}

type APITCPHealthCheck struct{}

// ToAPIRequest converts the Terraform model to an API request.
func (m *HealthcheckResourceModel) ToAPIRequest() *APIHealthcheck {
	spec := APIHealthcheckSpec{}

	if !m.Timeout.IsNull() && !m.Timeout.IsUnknown() {
		spec.Timeout = uint32(m.Timeout.ValueInt64())
	}
	if !m.Interval.IsNull() && !m.Interval.IsUnknown() {
		spec.Interval = uint32(m.Interval.ValueInt64())
	}
	if !m.UnhealthyThreshold.IsNull() && !m.UnhealthyThreshold.IsUnknown() {
		spec.UnhealthyThreshold = uint32(m.UnhealthyThreshold.ValueInt64())
	}
	if !m.HealthyThreshold.IsNull() && !m.HealthyThreshold.IsUnknown() {
		spec.HealthyThreshold = uint32(m.HealthyThreshold.ValueInt64())
	}

	if m.HTTPHealthCheck != nil {
		httpCheck := &APIHTTPHealthCheck{
			Path: m.HTTPHealthCheck.Path.ValueString(),
		}
		if m.HTTPHealthCheck.UseHTTPS.ValueBool() {
			httpCheck.UseHTTPS = &struct{}{}
		}
		if len(m.HTTPHealthCheck.ExpectedStatusCodes) > 0 {
			codes := make([]string, len(m.HTTPHealthCheck.ExpectedStatusCodes))
			for i, c := range m.HTTPHealthCheck.ExpectedStatusCodes {
				codes[i] = c.ValueString()
			}
			httpCheck.ExpectedStatusCodes = codes
		}
		spec.HTTPHealthCheck = httpCheck
	}

	if m.TCPHealthCheck != nil {
		spec.TCPHealthCheck = &APITCPHealthCheck{}
	}

	return &APIHealthcheck{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

// FromAPIResponse updates the model from an API response.
func (m *HealthcheckResourceModel) FromAPIResponse(resp *APIHealthcheck) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	if resp.Spec.Timeout > 0 {
		m.Timeout = types.Int64Value(int64(resp.Spec.Timeout))
	}
	if resp.Spec.Interval > 0 {
		m.Interval = types.Int64Value(int64(resp.Spec.Interval))
	}
	if resp.Spec.UnhealthyThreshold > 0 {
		m.UnhealthyThreshold = types.Int64Value(int64(resp.Spec.UnhealthyThreshold))
	}
	if resp.Spec.HealthyThreshold > 0 {
		m.HealthyThreshold = types.Int64Value(int64(resp.Spec.HealthyThreshold))
	}

	if resp.Spec.HTTPHealthCheck != nil {
		m.HTTPHealthCheck = &HTTPHealthCheckModel{
			Path:     types.StringValue(resp.Spec.HTTPHealthCheck.Path),
			UseHTTPS: types.BoolValue(resp.Spec.HTTPHealthCheck.UseHTTPS != nil),
		}
		if len(resp.Spec.HTTPHealthCheck.ExpectedStatusCodes) > 0 {
			codes := make([]types.String, len(resp.Spec.HTTPHealthCheck.ExpectedStatusCodes))
			for i, c := range resp.Spec.HTTPHealthCheck.ExpectedStatusCodes {
				codes[i] = types.StringValue(c)
			}
			m.HTTPHealthCheck.ExpectedStatusCodes = codes
		}
	}

	if resp.Spec.TCPHealthCheck != nil {
		m.TCPHealthCheck = &TCPHealthCheckModel{}
	}

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
