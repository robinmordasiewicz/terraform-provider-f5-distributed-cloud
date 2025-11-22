// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_domain

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ resource.Resource = &DNSDomainResource{}
var _ resource.ResourceWithImportState = &DNSDomainResource{}

type DNSDomainResource struct {
	client *client.Client
}

func NewDNSDomainResource() resource.Resource { return &DNSDomainResource{} }

func (r *DNSDomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_domain"
}

func (r *DNSDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

func (r *DNSDomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DNSDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DNSDomainResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Creating DNS Domain", map[string]interface{}{"name": data.Name.ValueString()})
	apiReq := data.ToAPIRequest()
	var apiResp APIDNSDomain
	path := fmt.Sprintf("/config/namespaces/%s/dns_domains", data.Namespace.ValueString())
	if err := r.client.Post(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error creating DNS Domain", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DNSDomainResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var apiResp APIDNSDomain
	path := fmt.Sprintf("/config/namespaces/%s/dns_domains/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading DNS Domain", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSDomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DNSDomainResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updating DNS Domain", map[string]interface{}{"name": data.Name.ValueString()})
	apiReq := data.ToAPIRequest()
	var apiResp APIDNSDomain
	path := fmt.Sprintf("/config/namespaces/%s/dns_domains/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Put(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error updating DNS Domain", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSDomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DNSDomainResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Deleting DNS Domain", map[string]interface{}{"name": data.Name.ValueString()})
	path := fmt.Sprintf("/config/namespaces/%s/dns_domains/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Delete(ctx, path); err != nil {
		resp.Diagnostics.AddError("Error deleting DNS Domain", err.Error())
		return
	}
}

func (r *DNSDomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid import ID", "Expected format: namespace/name")
		return
	}
	var data DNSDomainResourceModel
	data.Namespace = types.StringValue(parts[0])
	data.Name = types.StringValue(parts[1])
	var apiResp APIDNSDomain
	path := fmt.Sprintf("/config/namespaces/%s/dns_domains/%s", parts[0], parts[1])
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error importing DNS Domain", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
