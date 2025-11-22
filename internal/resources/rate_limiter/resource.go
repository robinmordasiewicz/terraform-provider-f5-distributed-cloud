// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package rate_limiter

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ resource.Resource = &RateLimiterResource{}
var _ resource.ResourceWithImportState = &RateLimiterResource{}

type RateLimiterResource struct {
	client *client.Client
}

func NewRateLimiterResource() resource.Resource { return &RateLimiterResource{} }

func (r *RateLimiterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rate_limiter"
}

func (r *RateLimiterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

func (r *RateLimiterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RateLimiterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RateLimiterResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Creating Rate Limiter", map[string]interface{}{"name": data.Name.ValueString()})
	apiReq := data.ToAPIRequest()
	var apiResp APIRateLimiter
	path := fmt.Sprintf("/config/namespaces/%s/rate_limiters", data.Namespace.ValueString())
	if err := r.client.Post(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error creating Rate Limiter", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RateLimiterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RateLimiterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var apiResp APIRateLimiter
	path := fmt.Sprintf("/config/namespaces/%s/rate_limiters/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading Rate Limiter", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RateLimiterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RateLimiterResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	apiReq := data.ToAPIRequest()
	var apiResp APIRateLimiter
	path := fmt.Sprintf("/config/namespaces/%s/rate_limiters/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Put(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error updating Rate Limiter", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RateLimiterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RateLimiterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	path := fmt.Sprintf("/config/namespaces/%s/rate_limiters/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Delete(ctx, path); err != nil && !client.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting Rate Limiter", err.Error())
	}
}

func (r *RateLimiterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid import ID", fmt.Sprintf("Expected namespace/name, got: %s", req.ID))
		return
	}
	var apiResp APIRateLimiter
	path := fmt.Sprintf("/config/namespaces/%s/rate_limiters/%s", parts[0], parts[1])
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error importing Rate Limiter", err.Error())
		return
	}
	var data RateLimiterResourceModel
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
