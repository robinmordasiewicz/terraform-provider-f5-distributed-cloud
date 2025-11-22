// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package user

import "github.com/hashicorp/terraform-plugin-framework/types"

type UserResourceModel struct {
	Name        types.String       `tfsdk:"name"`
	Namespace   types.String       `tfsdk:"namespace"`
	Email       types.String       `tfsdk:"email"`
	FirstName   types.String       `tfsdk:"first_name"`
	LastName    types.String       `tfsdk:"last_name"`
	Type        types.String       `tfsdk:"type"`
	Roles       []RoleRefModel     `tfsdk:"roles"`
	ID          types.String       `tfsdk:"id"`
}

type RoleRefModel struct {
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
}

type APIUser struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIUserSpec       `json:"spec"`
	SystemMeta APISystemMetadata `json:"system_metadata,omitempty"`
}

type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
}

type APISystemMetadata struct {
	UID               string `json:"uid,omitempty"`
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
}

type APIUserSpec struct {
	Email     string       `json:"email,omitempty"`
	FirstName string       `json:"first_name,omitempty"`
	LastName  string       `json:"last_name,omitempty"`
	Type      string       `json:"type,omitempty"`
	Roles     []APIRoleRef `json:"namespace_roles,omitempty"`
}

type APIRoleRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

func (m *UserResourceModel) ToAPIRequest() *APIUser {
	spec := APIUserSpec{
		Email:     m.Email.ValueString(),
		FirstName: m.FirstName.ValueString(),
		LastName:  m.LastName.ValueString(),
		Type:      m.Type.ValueString(),
	}
	if len(m.Roles) > 0 {
		spec.Roles = make([]APIRoleRef, len(m.Roles))
		for i, r := range m.Roles {
			spec.Roles[i] = APIRoleRef{
				Name:      r.Name.ValueString(),
				Namespace: r.Namespace.ValueString(),
			}
		}
	}
	return &APIUser{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString()},
		Spec:     spec,
	}
}

func (m *UserResourceModel) FromAPIResponse(resp *APIUser) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Email = types.StringValue(resp.Spec.Email)
	m.FirstName = types.StringValue(resp.Spec.FirstName)
	m.LastName = types.StringValue(resp.Spec.LastName)
	m.Type = types.StringValue(resp.Spec.Type)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
