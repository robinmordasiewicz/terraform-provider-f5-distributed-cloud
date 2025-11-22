// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package policy_based_routing

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PolicyBasedRoutingResourceModel describes the resource data model.
type PolicyBasedRoutingResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

// APIPolicyBasedRouting represents the API request/response structure.
type APIPolicyBasedRouting struct {
	Metadata   APIMetadata                 `json:"metadata"`
	Spec       APIPolicyBasedRoutingSpec   `json:"spec"`
	SystemMeta *APISystemMetadata          `json:"system_metadata,omitempty"`
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

// APIPolicyBasedRoutingSpec represents the Policy Based Routing specification.
type APIPolicyBasedRoutingSpec struct {
}

// ToAPIRequest converts the Terraform model to API request format.
func (m *PolicyBasedRoutingResourceModel) ToAPIRequest() *APIPolicyBasedRouting {
	apiReq := &APIPolicyBasedRouting{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIPolicyBasedRoutingSpec{},
	}

	return apiReq
}

// FromAPIResponse updates the Terraform model from API response.
func (m *PolicyBasedRoutingResourceModel) FromAPIResponse(resp *APIPolicyBasedRouting) {
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
