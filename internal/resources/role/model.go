// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package role

import "github.com/hashicorp/terraform-plugin-framework/types"

type RoleResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	RoleType    types.String `tfsdk:"role_type"`
	Permissions []PermissionModel `tfsdk:"permissions"`
	ID          types.String `tfsdk:"id"`
}

type PermissionModel struct {
	Resource   types.String `tfsdk:"resource"`
	Operations types.List   `tfsdk:"operations"`
}

type APIRole struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIRoleSpec       `json:"spec"`
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

type APIRoleSpec struct {
	RoleType    string          `json:"role_type,omitempty"`
	Permissions []APIPermission `json:"permissions,omitempty"`
}

type APIPermission struct {
	Resource   string   `json:"resource,omitempty"`
	Operations []string `json:"operations,omitempty"`
}

func (m *RoleResourceModel) ToAPIRequest() *APIRole {
	spec := APIRoleSpec{}
	if !m.RoleType.IsNull() {
		spec.RoleType = m.RoleType.ValueString()
	}
	if len(m.Permissions) > 0 {
		spec.Permissions = make([]APIPermission, len(m.Permissions))
		for i, perm := range m.Permissions {
			apiPerm := APIPermission{
				Resource: perm.Resource.ValueString(),
			}
			if !perm.Operations.IsNull() {
				ops := make([]string, 0)
				perm.Operations.ElementsAs(nil, &ops, false)
				apiPerm.Operations = ops
			}
			spec.Permissions[i] = apiPerm
		}
	}
	return &APIRole{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *RoleResourceModel) FromAPIResponse(resp *APIRole) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.Spec.RoleType != "" {
		m.RoleType = types.StringValue(resp.Spec.RoleType)
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
