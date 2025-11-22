// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package crl

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// CRLResourceModel describes the resource data model.
type CRLResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
}

// APICRL represents the API request/response structure.
type APICRL struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APICRLSpec         `json:"spec"`
	SystemMeta *APISystemMetadata `json:"system_metadata,omitempty"`
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

// APICRLSpec represents the CRL specification.
type APICRLSpec struct {
}

// ToAPIRequest converts the Terraform model to API request format.
func (m *CRLResourceModel) ToAPIRequest() *APICRL {
	apiReq := &APICRL{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APICRLSpec{},
	}

	return apiReq
}

// FromAPIResponse updates the Terraform model from API response.
func (m *CRLResourceModel) FromAPIResponse(resp *APICRL) {
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
