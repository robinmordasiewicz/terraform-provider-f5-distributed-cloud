// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package bgp

import "github.com/hashicorp/terraform-plugin-framework/types"

type BGPResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	ASN         types.Int64  `tfsdk:"asn"`
	ID          types.String `tfsdk:"id"`
}

type APIBGP struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIBGPSpec        `json:"spec"`
	SystemMeta APISystemMetadata `json:"system_metadata,omitempty"`
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

type APIBGPSpec struct {
	ASN int64 `json:"asn,omitempty"`
}

func (m *BGPResourceModel) ToAPIRequest() *APIBGP {
	spec := APIBGPSpec{}
	if !m.ASN.IsNull() {
		spec.ASN = m.ASN.ValueInt64()
	}
	return &APIBGP{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *BGPResourceModel) FromAPIResponse(resp *APIBGP) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
	if resp.Spec.ASN != 0 {
		m.ASN = types.Int64Value(resp.Spec.ASN)
	}
}
