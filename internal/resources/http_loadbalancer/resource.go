// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package http_loadbalancer

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

// Ensure provider-defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &HTTPLoadBalancerResource{}
	_ resource.ResourceWithImportState = &HTTPLoadBalancerResource{}
)

// NewHTTPLoadBalancerResource creates a new HTTP Load Balancer resource.
func NewHTTPLoadBalancerResource() resource.Resource {
	return &HTTPLoadBalancerResource{}
}

// HTTPLoadBalancerResource defines the resource implementation.
type HTTPLoadBalancerResource struct {
	client *client.Client
}

// Metadata returns the resource type name.
func (r *HTTPLoadBalancerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_http_loadbalancer"
}

// Schema returns the resource schema.
func (r *HTTPLoadBalancerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

// Configure adds the provider configured client to the resource.
func (r *HTTPLoadBalancerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T", req.ProviderData),
		)
		return
	}

	r.client = c
}

// Create creates the HTTP Load Balancer resource.
func (r *HTTPLoadBalancerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HTTPLoadBalancerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating HTTP Load Balancer", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Build API request
	apiReq := data.ToAPIRequest()

	// Create HTTP Load Balancer via API
	var apiResp APIHTTPLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/http_loadbalancers", data.Namespace.ValueString())
	err := r.client.Post(ctx, path, apiReq, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating HTTP Load Balancer",
			fmt.Sprintf("Could not create HTTP Load Balancer %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	tflog.Debug(ctx, "Created HTTP Load Balancer", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
		"id":        data.ID.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read reads the HTTP Load Balancer resource.
func (r *HTTPLoadBalancerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HTTPLoadBalancerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading HTTP Load Balancer", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Read HTTP Load Balancer from API
	var apiResp APIHTTPLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/http_loadbalancers/%s",
		data.Namespace.ValueString(), data.Name.ValueString())
	err := r.client.Get(ctx, path, &apiResp)
	if err != nil {
		if client.IsNotFound(err) {
			tflog.Debug(ctx, "HTTP Load Balancer not found, removing from state", map[string]interface{}{
				"name":      data.Name.ValueString(),
				"namespace": data.Namespace.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error Reading HTTP Load Balancer",
			fmt.Sprintf("Could not read HTTP Load Balancer %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update updates the HTTP Load Balancer resource.
func (r *HTTPLoadBalancerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data HTTPLoadBalancerResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating HTTP Load Balancer", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Build API request
	apiReq := data.ToAPIRequest()

	// Update HTTP Load Balancer via API
	var apiResp APIHTTPLoadBalancer
	path := fmt.Sprintf("/config/namespaces/%s/http_loadbalancers/%s",
		data.Namespace.ValueString(), data.Name.ValueString())
	err := r.client.Put(ctx, path, apiReq, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating HTTP Load Balancer",
			fmt.Sprintf("Could not update HTTP Load Balancer %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete deletes the HTTP Load Balancer resource.
func (r *HTTPLoadBalancerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HTTPLoadBalancerResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting HTTP Load Balancer", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Delete HTTP Load Balancer via API
	path := fmt.Sprintf("/config/namespaces/%s/http_loadbalancers/%s",
		data.Namespace.ValueString(), data.Name.ValueString())
	err := r.client.Delete(ctx, path)
	if err != nil {
		if client.IsNotFound(err) {
			tflog.Debug(ctx, "HTTP Load Balancer already deleted", map[string]interface{}{
				"name":      data.Name.ValueString(),
				"namespace": data.Namespace.ValueString(),
			})
			return
		}
		resp.Diagnostics.AddError(
			"Error Deleting HTTP Load Balancer",
			fmt.Sprintf("Could not delete HTTP Load Balancer %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err),
		)
		return
	}

	tflog.Debug(ctx, "Deleted HTTP Load Balancer", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})
}

// ImportState imports an existing HTTP Load Balancer into Terraform state.
func (r *HTTPLoadBalancerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Importing HTTP Load Balancer", map[string]interface{}{
		"id": req.ID,
	})

	// Expected format: namespace/name
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("Expected import ID in format 'namespace/name', got: %s", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("namespace"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), parts[1])...)
}
