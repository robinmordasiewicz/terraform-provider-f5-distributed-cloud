// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// MetadataModel represents the common metadata block for F5 XC resources.
type MetadataModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Labels      types.Map    `tfsdk:"labels"`
	Annotations types.Map    `tfsdk:"annotations"`
	Description types.String `tfsdk:"description"`
	Disable     types.Bool   `tfsdk:"disable"`
}

// MetadataSchema returns the schema for the metadata block.
func MetadataSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: "Metadata for the resource.",
		Required:    true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "Name of the resource. Must be unique within the namespace.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"namespace": schema.StringAttribute{
				Description: "Namespace where the resource is created.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"labels": schema.MapAttribute{
				Description: "Labels for the resource.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"annotations": schema.MapAttribute{
				Description: "Annotations for the resource.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"description": schema.StringAttribute{
				Description: "Description of the resource.",
				Optional:    true,
			},
			"disable": schema.BoolAttribute{
				Description: "Whether the resource is disabled.",
				Optional:    true,
			},
		},
	}
}

// SystemMetadataModel represents the system-managed metadata.
type SystemMetadataModel struct {
	UID               types.String `tfsdk:"uid"`
	CreationTimestamp types.String `tfsdk:"creation_timestamp"`
	ModificationTime  types.String `tfsdk:"modification_time"`
	CreatorClass      types.String `tfsdk:"creator_class"`
	CreatorID         types.String `tfsdk:"creator_id"`
}

// SystemMetadataSchema returns the schema for system metadata (computed).
func SystemMetadataSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Description: "System-managed metadata.",
		Computed:    true,
		Attributes: map[string]schema.Attribute{
			"uid": schema.StringAttribute{
				Description: "Unique identifier of the resource.",
				Computed:    true,
			},
			"creation_timestamp": schema.StringAttribute{
				Description: "Timestamp when the resource was created.",
				Computed:    true,
			},
			"modification_time": schema.StringAttribute{
				Description: "Timestamp when the resource was last modified.",
				Computed:    true,
			},
			"creator_class": schema.StringAttribute{
				Description: "Class of the creator.",
				Computed:    true,
			},
			"creator_id": schema.StringAttribute{
				Description: "ID of the creator.",
				Computed:    true,
			},
		},
	}
}

// SystemMetadataType returns the object type for system metadata.
func SystemMetadataType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"uid":                types.StringType,
			"creation_timestamp": types.StringType,
			"modification_time":  types.StringType,
			"creator_class":      types.StringType,
			"creator_id":         types.StringType,
		},
	}
}

// SpecMetadataModel represents the common spec metadata.
type SpecMetadataModel struct {
	GCSpec types.Object `tfsdk:"gc_spec"`
}

// APIMetadata represents metadata in API requests/responses.
type APIMetadata struct {
	Name        string            `json:"name,omitempty"`
	Namespace   string            `json:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	Description string            `json:"description,omitempty"`
	Disable     bool              `json:"disable,omitempty"`
}

// APISystemMetadata represents system metadata from API responses.
type APISystemMetadata struct {
	UID               string `json:"uid,omitempty"`
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
	ModificationTime  string `json:"modification_time,omitempty"`
	CreatorClass      string `json:"creator_class,omitempty"`
	CreatorID         string `json:"creator_id,omitempty"`
}
