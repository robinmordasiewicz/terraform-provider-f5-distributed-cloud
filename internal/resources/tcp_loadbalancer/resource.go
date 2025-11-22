// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package tcp_loadbalancer

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TCPLoadBalancerResource{}
var _ resource.ResourceWithImportState = &TCPLoadBalancerResource{}

// TCPLoadBalancerResource defines the resource implementation.
type TCPLoadBalancerResource struct {
	client *client.Client
}

// NewTCPLoadBalancerResource creates a new resource instance.
func NewTCPLoadBalancerResource() resource.Resource {
	return &TCPLoadBalancerResource{}
}

// Metadata returns the resource type name.
func (r *TCPLoadBalancerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tcp_loadbalancer"
}

// Schema returns the resource schema.
func (r *TCPLoadBalancerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

// Configure configures the resource with provider data.
func (r *TCPLoadBalancerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = c
}

// Create creates the resource and sets the initial Terraform state.
func (r *TCPLoadBalancerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TCPLoadBalancerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating TCP Load Balancer", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Build API request
	apiReq := data.ToAPIRequest()

	// Call API
	var apiResp APITCPLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/tcp_loadbalancers", data.Namespace.ValueString())
	if err := r.client.Post(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError(
			"Error creating TCP Load Balancer",
			fmt.Sprintf("Could not create TCP Load Balancer, unexpected error: %s", err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	tflog.Trace(ctx, "Created TCP Load Balancer", map[string]interface{}{
		"id": data.ID.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read reads the resource state.
func (r *TCPLoadBalancerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TCPLoadBalancerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading TCP Load Balancer", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Call API
	var apiResp APITCPLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/tcp_loadbalancers/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading TCP Load Balancer",
			fmt.Sprintf("Could not read TCP Load Balancer, unexpected error: %s", err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update updates the resource.
func (r *TCPLoadBalancerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TCPLoadBalancerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating TCP Load Balancer", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Build API request
	apiReq := data.ToAPIRequest()

	// Call API
	var apiResp APITCPLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/tcp_loadbalancers/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Put(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError(
			"Error updating TCP Load Balancer",
			fmt.Sprintf("Could not update TCP Load Balancer, unexpected error: %s", err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete deletes the resource.
func (r *TCPLoadBalancerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TCPLoadBalancerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting TCP Load Balancer", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Call API
	path := fmt.Sprintf("/config/namespaces/%s/tcp_loadbalancers/%s", data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Delete(ctx, path); err != nil {
		if client.IsNotFound(err) {
			return
		}
		resp.Diagnostics.AddError(
			"Error deleting TCP Load Balancer",
			fmt.Sprintf("Could not delete TCP Load Balancer, unexpected error: %s", err),
		)
		return
	}
}

// ImportState imports the resource state.
func (r *TCPLoadBalancerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: namespace/name
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			fmt.Sprintf("Expected import ID format: namespace/name, got: %s", req.ID),
		)
		return
	}

	namespace := parts[0]
	name := parts[1]

	tflog.Debug(ctx, "Importing TCP Load Balancer", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})

	// Read from API
	var apiResp APITCPLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/tcp_loadbalancers/%s", namespace, name)
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		resp.Diagnostics.AddError(
			"Error importing TCP Load Balancer",
			fmt.Sprintf("Could not import TCP Load Balancer, unexpected error: %s", err),
		)
		return
	}

	// Create model from response
	var data TCPLoadBalancerResourceModel
	data.FromAPIResponse(&apiResp)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
