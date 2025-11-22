// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package k8s_cluster_role

import "github.com/hashicorp/terraform-plugin-framework/types"

type K8sClusterRoleResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
}

type APIK8sClusterRole struct {
	Metadata   APIMetadata           `json:"metadata"`
	Spec       APIK8sClusterRoleSpec `json:"spec"`
	SystemMeta APISystemMetadata     `json:"system_metadata,omitempty"`
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

type APIK8sClusterRoleSpec struct {
	// K8s cluster role configuration fields
}

func (m *K8sClusterRoleResourceModel) ToAPIRequest() *APIK8sClusterRole {
	return &APIK8sClusterRole{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     APIK8sClusterRoleSpec{},
	}
}

func (m *K8sClusterRoleResourceModel) FromAPIResponse(resp *APIK8sClusterRole) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
