// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_lb_health_check

import "github.com/hashicorp/terraform-plugin-framework/types"

type DNSLBHealthCheckResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Protocol    types.String `tfsdk:"protocol"`
	Port        types.Int64  `tfsdk:"port"`
	Interval    types.Int64  `tfsdk:"interval"`
	Timeout     types.Int64  `tfsdk:"timeout"`
	ID          types.String `tfsdk:"id"`
}

type APIDNSLBHealthCheck struct {
	Metadata   APIMetadata             `json:"metadata"`
	Spec       APIDNSLBHealthCheckSpec `json:"spec"`
	SystemMeta APISystemMetadata       `json:"system_metadata,omitempty"`
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

type APIDNSLBHealthCheckSpec struct {
	Protocol string `json:"protocol,omitempty"`
	Port     int64  `json:"port,omitempty"`
	Interval int64  `json:"interval,omitempty"`
	Timeout  int64  `json:"timeout,omitempty"`
}

func (m *DNSLBHealthCheckResourceModel) ToAPIRequest() *APIDNSLBHealthCheck {
	spec := APIDNSLBHealthCheckSpec{
		Protocol: m.Protocol.ValueString(),
	}
	if !m.Port.IsNull() {
		spec.Port = m.Port.ValueInt64()
	}
	if !m.Interval.IsNull() {
		spec.Interval = m.Interval.ValueInt64()
	}
	if !m.Timeout.IsNull() {
		spec.Timeout = m.Timeout.ValueInt64()
	}
	return &APIDNSLBHealthCheck{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *DNSLBHealthCheckResourceModel) FromAPIResponse(resp *APIDNSLBHealthCheck) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.Protocol = types.StringValue(resp.Spec.Protocol)
	if resp.Spec.Port != 0 {
		m.Port = types.Int64Value(resp.Spec.Port)
	}
	if resp.Spec.Interval != 0 {
		m.Interval = types.Int64Value(resp.Spec.Interval)
	}
	if resp.Spec.Timeout != 0 {
		m.Timeout = types.Int64Value(resp.Spec.Timeout)
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
