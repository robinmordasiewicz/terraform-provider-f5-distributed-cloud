// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package k8s_cluster_role_binding

import "github.com/hashicorp/terraform-plugin-framework/types"

type K8sClusterRoleBindingResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
}

type APIK8sClusterRoleBinding struct {
	Metadata   APIMetadata                  `json:"metadata"`
	Spec       APIK8sClusterRoleBindingSpec `json:"spec"`
	SystemMeta APISystemMetadata            `json:"system_metadata,omitempty"`
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

type APIK8sClusterRoleBindingSpec struct {
	// K8s cluster role binding configuration fields
}

func (m *K8sClusterRoleBindingResourceModel) ToAPIRequest() *APIK8sClusterRoleBinding {
	return &APIK8sClusterRoleBinding{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     APIK8sClusterRoleBindingSpec{},
	}
}

func (m *K8sClusterRoleBindingResourceModel) FromAPIResponse(resp *APIK8sClusterRoleBinding) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
