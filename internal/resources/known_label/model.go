// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package known_label

import "github.com/hashicorp/terraform-plugin-framework/types"

type KnownLabelResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Key         types.String `tfsdk:"key"`
	Value       types.String `tfsdk:"value"`
	ID          types.String `tfsdk:"id"`
}

type APIKnownLabel struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIKnownLabelSpec `json:"spec"`
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

type APIKnownLabelSpec struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func (m *KnownLabelResourceModel) ToAPIRequest() *APIKnownLabel {
	return &APIKnownLabel{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIKnownLabelSpec{
			Key:   m.Key.ValueString(),
			Value: m.Value.ValueString(),
		},
	}
}

func (m *KnownLabelResourceModel) FromAPIResponse(resp *APIKnownLabel) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.Key = types.StringValue(resp.Spec.Key)
	m.Value = types.StringValue(resp.Spec.Value)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
