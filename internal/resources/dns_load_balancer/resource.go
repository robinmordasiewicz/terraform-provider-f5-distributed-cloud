// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package dns_load_balancer

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var _ resource.Resource = &DNSLoadBalancerResource{}
var _ resource.ResourceWithImportState = &DNSLoadBalancerResource{}

type DNSLoadBalancerResource struct {
	client *client.Client
}

func NewDNSLoadBalancerResource() resource.Resource { return &DNSLoadBalancerResource{} }

func (r *DNSLoadBalancerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_load_balancer"
}

func (r *DNSLoadBalancerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

func (r *DNSLoadBalancerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DNSLoadBalancerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DNSLoadBalancerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Creating DNS Load Balancer", map[string]interface{}{"name": data.Name.ValueString()})
	apiReq := data.ToAPIRequest()
	var apiResp APIDNSLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/dns_load_balancers", data.Namespace.ValueString())
	if err := r.client.Post(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error creating DNS Load Balancer", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSLoadBalancerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DNSLoadBalancerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var apiResp APIDNSLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/dns_load_balancers/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading DNS Load Balancer", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSLoadBalancerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DNSLoadBalancerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Updating DNS Load Balancer", map[string]interface{}{"name": data.Name.ValueString()})
	apiReq := data.ToAPIRequest()
	var apiResp APIDNSLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/dns_load_balancers/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Put(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error updating DNS Load Balancer", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DNSLoadBalancerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DNSLoadBalancerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Deleting DNS Load Balancer", map[string]interface{}{"name": data.Name.ValueString()})
	path := fmt.Sprintf("/config/namespaces/%s/dns_load_balancers/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Delete(ctx, path); err != nil {
		resp.Diagnostics.AddError("Error deleting DNS Load Balancer", err.Error())
		return
	}
}

func (r *DNSLoadBalancerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid import ID", "Expected format: namespace/name")
		return
	}
	var data DNSLoadBalancerResourceModel
	data.Namespace = types.StringValue(parts[0])
	data.Name = types.StringValue(parts[1])
	var apiResp APIDNSLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/dns_load_balancers/%s", parts[0], parts[1])
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error importing DNS Load Balancer", err.Error())
		return
	}
	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
