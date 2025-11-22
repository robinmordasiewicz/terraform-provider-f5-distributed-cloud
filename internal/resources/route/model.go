// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RouteResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	Path        types.String `tfsdk:"path"`
	Methods     types.List   `tfsdk:"methods"`
	Destination types.String `tfsdk:"destination"`
	ID          types.String `tfsdk:"id"`
}

type APIRoute struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIRouteSpec      `json:"spec"`
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

type APIRouteSpec struct {
	Path        string   `json:"path,omitempty"`
	Methods     []string `json:"methods,omitempty"`
	Destination string   `json:"destination,omitempty"`
}

func (m *RouteResourceModel) ToAPIRequest(ctx context.Context) *APIRoute {
	spec := APIRouteSpec{
		Path:        m.Path.ValueString(),
		Destination: m.Destination.ValueString(),
	}
	if !m.Methods.IsNull() {
		var methods []string
		m.Methods.ElementsAs(ctx, &methods, false)
		spec.Methods = methods
	}
	return &APIRoute{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

func (m *RouteResourceModel) FromAPIResponse(ctx context.Context, resp *APIRoute) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.Path = types.StringValue(resp.Spec.Path)
	m.Destination = types.StringValue(resp.Spec.Destination)
	if len(resp.Spec.Methods) > 0 {
		methods, _ := types.ListValueFrom(ctx, types.StringType, resp.Spec.Methods)
		m.Methods = methods
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
