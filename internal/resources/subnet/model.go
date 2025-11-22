// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package subnet

import "github.com/hashicorp/terraform-plugin-framework/types"

type SubnetResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	CIDR        types.String `tfsdk:"cidr"`
	ID          types.String `tfsdk:"id"`
}

type APISubnet struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APISubnetSpec     `json:"spec"`
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

type APISubnetSpec struct {
	CIDR string `json:"cidr,omitempty"`
}

func (m *SubnetResourceModel) ToAPIRequest() *APISubnet {
	return &APISubnet{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     APISubnetSpec{CIDR: m.CIDR.ValueString()},
	}
}

func (m *SubnetResourceModel) FromAPIResponse(resp *APISubnet) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.CIDR = types.StringValue(resp.Spec.CIDR)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
