// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package bgp_asn_set

import "github.com/hashicorp/terraform-plugin-framework/types"

type BGPASNSetResourceModel struct {
	Name        types.String  `tfsdk:"name"`
	Namespace   types.String  `tfsdk:"namespace"`
	Description types.String  `tfsdk:"description"`
	ASNs        []types.Int64 `tfsdk:"asns"`
	ID          types.String  `tfsdk:"id"`
}

type APIBGPASNSet struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIBGPASNSetSpec  `json:"spec"`
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

type APIBGPASNSetSpec struct {
	ASNs []int64 `json:"asn_list,omitempty"`
}

func (m *BGPASNSetResourceModel) ToAPIRequest() *APIBGPASNSet {
	spec := APIBGPASNSetSpec{}
	if len(m.ASNs) > 0 {
		asns := make([]int64, len(m.ASNs))
		for i, asn := range m.ASNs {
			asns[i] = asn.ValueInt64()
		}
		spec.ASNs = asns
	}
	return &APIBGPASNSet{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *BGPASNSetResourceModel) FromAPIResponse(resp *APIBGPASNSet) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
	if len(resp.Spec.ASNs) > 0 {
		asns := make([]types.Int64, len(resp.Spec.ASNs))
		for i, asn := range resp.Spec.ASNs {
			asns[i] = types.Int64Value(asn)
		}
		m.ASNs = asns
	}
}
