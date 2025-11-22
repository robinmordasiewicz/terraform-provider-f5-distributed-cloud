// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package bgp_routing_policy

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var (
	_ resource.Resource                = &BGPRoutingPolicyResource{}
	_ resource.ResourceWithImportState = &BGPRoutingPolicyResource{}
)

func NewBGPRoutingPolicyResource() resource.Resource {
	return &BGPRoutingPolicyResource{}
}

type BGPRoutingPolicyResource struct {
	client *client.Client
}

func (r *BGPRoutingPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_routing_policy"
}

func (r *BGPRoutingPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

func (r *BGPRoutingPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData))
		return
	}
	r.client = c
}

func (r *BGPRoutingPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BGPRoutingPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating BGP Routing Policy", map[string]interface{}{
		"name": data.Name.ValueString(), "namespace": data.Namespace.ValueString(),
	})

	apiReq := data.ToAPIRequest()
	var apiResp APIBGPRoutingPolicy
	path := fmt.Sprintf("/config/namespaces/%s/bgp_routing_policys", data.Namespace.ValueString())
	if err := r.client.Post(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Creating BGP Routing Policy",
			fmt.Sprintf("Could not create BGP Routing Policy %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err))
		return
	}

	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BGPRoutingPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BGPRoutingPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIBGPRoutingPolicy
	path := fmt.Sprintf("/config/namespaces/%s/bgp_routing_policys/%s",
		data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading BGP Routing Policy",
			fmt.Sprintf("Could not read BGP Routing Policy %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err))
		return
	}

	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BGPRoutingPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BGPRoutingPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := data.ToAPIRequest()
	var apiResp APIBGPRoutingPolicy
	path := fmt.Sprintf("/config/namespaces/%s/bgp_routing_policys/%s",
		data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Put(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Updating BGP Routing Policy",
			fmt.Sprintf("Could not update BGP Routing Policy %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err))
		return
	}

	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BGPRoutingPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BGPRoutingPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := fmt.Sprintf("/config/namespaces/%s/bgp_routing_policys/%s",
		data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Delete(ctx, path); err != nil {
		if client.IsNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error Deleting BGP Routing Policy",
			fmt.Sprintf("Could not delete BGP Routing Policy %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err))
	}
}

func (r *BGPRoutingPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid Import ID",
			fmt.Sprintf("Expected format 'namespace/name', got: %s", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("namespace"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), parts[1])...)
}
