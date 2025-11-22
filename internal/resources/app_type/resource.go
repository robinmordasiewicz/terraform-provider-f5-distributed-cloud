// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package app_type

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ resource.Resource = &AppTypeResource{}
var _ resource.ResourceWithImportState = &AppTypeResource{}

type AppTypeResource struct {
	client *client.Client
}

func NewAppTypeResource() resource.Resource { return &AppTypeResource{} }

func (r *AppTypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_type"
}

func (r *AppTypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

func (r *AppTypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *client.Client, got: %T.", req.ProviderData))
		return
	}
	r.client = c
}

func (r *AppTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppTypeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Creating App Type", map[string]interface{}{"name": data.Name.ValueString()})
	apiReq := data.ToAPIRequest()
	var apiResp APIAppType
	path := fmt.Sprintf("/config/namespaces/%s/app_types", data.Namespace.ValueString())
	if err := r.client.Post(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error creating App Type", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var apiResp APIAppType
	path := fmt.Sprintf("/config/namespaces/%s/app_types/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading App Type", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppTypeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	apiReq := data.ToAPIRequest()
	var apiResp APIAppType
	path := fmt.Sprintf("/config/namespaces/%s/app_types/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Put(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error updating App Type", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	path := fmt.Sprintf("/config/namespaces/%s/app_types/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Delete(ctx, path); err != nil && !client.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting App Type", err.Error())
	}
}

func (r *AppTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid import ID", fmt.Sprintf("Expected namespace/name, got: %s", req.ID))
		return
	}
	var apiResp APIAppType
	path := fmt.Sprintf("/config/namespaces/%s/app_types/%s", parts[0], parts[1])
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error importing App Type", err.Error())
		return
	}
	var data AppTypeResourceModel
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
