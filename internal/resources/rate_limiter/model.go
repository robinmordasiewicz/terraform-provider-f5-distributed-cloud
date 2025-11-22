// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package rate_limiter

import "github.com/hashicorp/terraform-plugin-framework/types"

type RateLimiterResourceModel struct {
	Name            types.String `tfsdk:"name"`
	Namespace       types.String `tfsdk:"namespace"`
	Description     types.String `tfsdk:"description"`
	TotalNumber     types.Int64  `tfsdk:"total_number"`
	Unit            types.String `tfsdk:"unit"`
	BurstMultiplier types.Int64  `tfsdk:"burst_multiplier"`
	ID              types.String `tfsdk:"id"`
}

type APIRateLimiter struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APIRateLimiterSpec `json:"spec"`
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

type APIRateLimiterSpec struct {
	TotalNumber     int64  `json:"total_number,omitempty"`
	Unit            string `json:"unit,omitempty"`
	BurstMultiplier int64  `json:"burst_multiplier,omitempty"`
}

func (m *RateLimiterResourceModel) ToAPIRequest() *APIRateLimiter {
	return &APIRateLimiter{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIRateLimiterSpec{
			TotalNumber:     m.TotalNumber.ValueInt64(),
			Unit:            m.Unit.ValueString(),
			BurstMultiplier: m.BurstMultiplier.ValueInt64(),
		},
	}
}

func (m *RateLimiterResourceModel) FromAPIResponse(resp *APIRateLimiter) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.TotalNumber = types.Int64Value(resp.Spec.TotalNumber)
	m.Unit = types.StringValue(resp.Spec.Unit)
	m.BurstMultiplier = types.Int64Value(resp.Spec.BurstMultiplier)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
