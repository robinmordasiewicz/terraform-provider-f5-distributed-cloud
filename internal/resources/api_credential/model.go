// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package api_credential

import "github.com/hashicorp/terraform-plugin-framework/types"

type APICredentialResourceModel struct {
	Name                types.String `tfsdk:"name"`
	Namespace           types.String `tfsdk:"namespace"`
	Description         types.String `tfsdk:"description"`
	CredentialType      types.String `tfsdk:"credential_type"`
	ExpirationDays      types.Int64  `tfsdk:"expiration_days"`
	VirtualK8sName      types.String `tfsdk:"virtual_k8s_name"`
	VirtualK8sNamespace types.String `tfsdk:"virtual_k8s_namespace"`
	ID                  types.String `tfsdk:"id"`
	Data                types.String `tfsdk:"data"`
}

type APIAPICredential struct {
	Metadata   APIMetadata          `json:"metadata"`
	Spec       APIAPICredentialSpec `json:"spec"`
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

type APIAPICredentialSpec struct {
	Type                string `json:"type,omitempty"`
	ExpirationDays      int32  `json:"expiration_days,omitempty"`
	VirtualK8sName      string `json:"virtual_k8s_name,omitempty"`
	VirtualK8sNamespace string `json:"virtual_k8s_namespace,omitempty"`
}

func (m *APICredentialResourceModel) ToAPIRequest() *APIAPICredential {
	spec := APIAPICredentialSpec{}
	if !m.CredentialType.IsNull() {
		spec.Type = m.CredentialType.ValueString()
	}
	if !m.ExpirationDays.IsNull() {
		spec.ExpirationDays = int32(m.ExpirationDays.ValueInt64())
	}
	if !m.VirtualK8sName.IsNull() {
		spec.VirtualK8sName = m.VirtualK8sName.ValueString()
	}
	if !m.VirtualK8sNamespace.IsNull() {
		spec.VirtualK8sNamespace = m.VirtualK8sNamespace.ValueString()
	}
	return &APIAPICredential{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *APICredentialResourceModel) FromAPIResponse(resp *APIAPICredential) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.Spec.Type != "" {
		m.CredentialType = types.StringValue(resp.Spec.Type)
	}
	if resp.Spec.ExpirationDays > 0 {
		m.ExpirationDays = types.Int64Value(int64(resp.Spec.ExpirationDays))
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
