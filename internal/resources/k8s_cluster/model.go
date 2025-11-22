// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package k8s_cluster

import "github.com/hashicorp/terraform-plugin-framework/types"

type K8sClusterResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
}

type APIK8sCluster struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIK8sClusterSpec `json:"spec"`
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

type APIK8sClusterSpec struct {
	// K8s cluster configuration fields
}

func (m *K8sClusterResourceModel) ToAPIRequest() *APIK8sCluster {
	return &APIK8sCluster{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     APIK8sClusterSpec{},
	}
}

func (m *K8sClusterResourceModel) FromAPIResponse(resp *APIK8sCluster) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
