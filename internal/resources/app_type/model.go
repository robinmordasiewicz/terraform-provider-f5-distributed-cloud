// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package app_type

import "github.com/hashicorp/terraform-plugin-framework/types"

type AppTypeResourceModel struct {
	Name        types.String     `tfsdk:"name"`
	Namespace   types.String     `tfsdk:"namespace"`
	Description types.String     `tfsdk:"description"`
	Features    []FeatureModel   `tfsdk:"features"`
	ID          types.String     `tfsdk:"id"`
}

type FeatureModel struct {
	Type types.String `tfsdk:"type"`
}

type APIAppType struct {
	Metadata   APIMetadata      `json:"metadata"`
	Spec       APIAppTypeSpec   `json:"spec"`
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

type APIAppTypeSpec struct {
	Features []APIFeature `json:"features,omitempty"`
}

type APIFeature struct {
	Type string `json:"type,omitempty"`
}

func (m *AppTypeResourceModel) ToAPIRequest() *APIAppType {
	spec := APIAppTypeSpec{}
	if len(m.Features) > 0 {
		spec.Features = make([]APIFeature, len(m.Features))
		for i, f := range m.Features {
			spec.Features[i] = APIFeature{Type: f.Type.ValueString()}
		}
	}
	return &APIAppType{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *AppTypeResourceModel) FromAPIResponse(resp *APIAppType) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
