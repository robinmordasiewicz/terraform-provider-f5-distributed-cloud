// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package api_definition

import "github.com/hashicorp/terraform-plugin-framework/types"

type APIDefinitionResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	SwaggerSpec types.String `tfsdk:"swagger_spec"`
	OpenAPISpec types.String `tfsdk:"openapi_spec"`
	BasePath    types.String `tfsdk:"base_path"`
	ID          types.String `tfsdk:"id"`
}

type APIAPIDefinition struct {
	Metadata   APIMetadata          `json:"metadata"`
	Spec       APIAPIDefinitionSpec `json:"spec"`
	SystemMeta APISystemMetadata    `json:"system_metadata,omitempty"`
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

type APIAPIDefinitionSpec struct {
	SwaggerSpec string `json:"swagger_spec,omitempty"`
	OpenAPISpec string `json:"openapi_spec,omitempty"`
	BasePath    string `json:"base_path,omitempty"`
}

func (m *APIDefinitionResourceModel) ToAPIRequest() *APIAPIDefinition {
	spec := APIAPIDefinitionSpec{}
	if !m.SwaggerSpec.IsNull() {
		spec.SwaggerSpec = m.SwaggerSpec.ValueString()
	}
	if !m.OpenAPISpec.IsNull() {
		spec.OpenAPISpec = m.OpenAPISpec.ValueString()
	}
	if !m.BasePath.IsNull() {
		spec.BasePath = m.BasePath.ValueString()
	}
	return &APIAPIDefinition{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *APIDefinitionResourceModel) FromAPIResponse(resp *APIAPIDefinition) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.Spec.BasePath != "" {
		m.BasePath = types.StringValue(resp.Spec.BasePath)
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
