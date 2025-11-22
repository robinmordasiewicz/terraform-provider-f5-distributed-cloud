// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package bgp_routing_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// BGPRoutingPolicyResourceModel describes the resource data model.
type BGPRoutingPolicyResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

// APIBGPRoutingPolicy represents the API request/response structure.
type APIBGPRoutingPolicy struct {
	Metadata   APIMetadata             `json:"metadata"`
	Spec       APIBGPRoutingPolicySpec `json:"spec"`
	SystemMeta *APISystemMetadata      `json:"system_metadata,omitempty"`
}

// APIMetadata represents the metadata section.
type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}

// APISystemMetadata represents system-generated metadata.
type APISystemMetadata struct {
	UID string `json:"uid,omitempty"`
}

// APIBGPRoutingPolicySpec represents the BGP Routing Policy specification.
type APIBGPRoutingPolicySpec struct {
}

// ToAPIRequest converts the Terraform model to API request format.
func (m *BGPRoutingPolicyResourceModel) ToAPIRequest() *APIBGPRoutingPolicy {
	apiReq := &APIBGPRoutingPolicy{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIBGPRoutingPolicySpec{},
	}

	return apiReq
}

// FromAPIResponse updates the Terraform model from API response.
func (m *BGPRoutingPolicyResourceModel) FromAPIResponse(resp *APIBGPRoutingPolicy) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)

	if resp.Metadata.Description != "" {
		m.Description = types.StringValue(resp.Metadata.Description)
	}

	if resp.SystemMeta != nil && resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
