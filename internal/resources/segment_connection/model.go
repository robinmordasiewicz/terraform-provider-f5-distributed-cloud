// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package segment_connection

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SegmentConnectionResourceModel describes the resource data model.
type SegmentConnectionResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`

	// Segment connection configuration
	Segment1 *SegmentRefModel `tfsdk:"segment1"`
	Segment2 *SegmentRefModel `tfsdk:"segment2"`

	// Computed fields
	ID types.String `tfsdk:"id"`
}

// SegmentRefModel represents a reference to a network segment.
type SegmentRefModel struct {
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
}

// API structures

type APISegmentConnection struct {
	Metadata   APIMetadata              `json:"metadata"`
	Spec       APISegmentConnectionSpec `json:"spec"`
	SystemMeta APISystemMetadata        `json:"system_metadata,omitempty"`
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

type APISegmentConnectionSpec struct {
	Segment1 *APISegmentRef `json:"segment1,omitempty"`
	Segment2 *APISegmentRef `json:"segment2,omitempty"`
}

type APISegmentRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *SegmentConnectionResourceModel) ToAPIRequest() *APISegmentConnection {
	spec := APISegmentConnectionSpec{}

	if m.Segment1 != nil {
		spec.Segment1 = &APISegmentRef{
			Name:      m.Segment1.Name.ValueString(),
			Namespace: m.Segment1.Namespace.ValueString(),
		}
	}

	if m.Segment2 != nil {
		spec.Segment2 = &APISegmentRef{
			Name:      m.Segment2.Name.ValueString(),
			Namespace: m.Segment2.Namespace.ValueString(),
		}
	}

	return &APISegmentConnection{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

// FromAPIResponse updates the model from an API response.
func (m *SegmentConnectionResourceModel) FromAPIResponse(resp *APISegmentConnection) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	if resp.Spec.Segment1 != nil {
		m.Segment1 = &SegmentRefModel{
			Name:      types.StringValue(resp.Spec.Segment1.Name),
			Namespace: types.StringValue(resp.Spec.Segment1.Namespace),
		}
	}

	if resp.Spec.Segment2 != nil {
		m.Segment2 = &SegmentRefModel{
			Name:      types.StringValue(resp.Spec.Segment2.Name),
			Namespace: types.StringValue(resp.Spec.Segment2.Namespace),
		}
	}

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
