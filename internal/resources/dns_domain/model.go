// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_domain

import "github.com/hashicorp/terraform-plugin-framework/types"

type DNSDomainResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Domain      types.String `tfsdk:"domain"`
	ID          types.String `tfsdk:"id"`
}

type APIDNSDomain struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIDNSDomainSpec  `json:"spec"`
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

type APIDNSDomainSpec struct {
	Domain string `json:"domain,omitempty"`
}

func (m *DNSDomainResourceModel) ToAPIRequest() *APIDNSDomain {
	return &APIDNSDomain{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     APIDNSDomainSpec{Domain: m.Domain.ValueString()},
	}
}

func (m *DNSDomainResourceModel) FromAPIResponse(resp *APIDNSDomain) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.Domain = types.StringValue(resp.Spec.Domain)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
