// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package network_policy_set

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworkPolicySetResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

type APINetworkPolicySet struct {
	Metadata   APIMetadata             `json:"metadata"`
	Spec       APINetworkPolicySetSpec `json:"spec"`
	SystemMeta *APISystemMetadata      `json:"system_metadata,omitempty"`
}

type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}

type APISystemMetadata struct {
	UID string `json:"uid,omitempty"`
}

type APINetworkPolicySetSpec struct {
}

func (m *NetworkPolicySetResourceModel) ToAPIRequest() *APINetworkPolicySet {
	return &APINetworkPolicySet{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APINetworkPolicySetSpec{},
	}
}

func (m *NetworkPolicySetResourceModel) FromAPIResponse(resp *APINetworkPolicySet) {
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
