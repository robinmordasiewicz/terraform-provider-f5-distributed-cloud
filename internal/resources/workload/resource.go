// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package workload

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/client"
)

// Ensure provider-defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &WorkloadResource{}
	_ resource.ResourceWithImportState = &WorkloadResource{}
)

// NewWorkloadResource creates a new workload resource.
func NewWorkloadResource() resource.Resource {
	return &WorkloadResource{}
}

// WorkloadResource defines the resource implementation.
type WorkloadResource struct {
	client *client.Client
}

// Metadata returns the resource type name.
func (r *WorkloadResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workload"
}

// Schema returns the resource schema.
func (r *WorkloadResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

// Configure adds the provider configured client to the resource.
func (r *WorkloadResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T.", req.ProviderData),
		)
		return
	}

	r.client = c
}

// Create creates the workload resource.
func (r *WorkloadResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data WorkloadResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating Workload", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Build API request
	apiReq := data.ToAPIRequest()

	// Create workload via API
	var apiResp APIWorkload
	path := fmt.Sprintf("/config/namespaces/%s/workloads", data.Namespace.ValueString())
	err := r.client.Post(ctx, path, apiReq, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Workload",
			fmt.Sprintf("Could not create workload %s: %s", data.Name.ValueString(), err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	tflog.Debug(ctx, "Created Workload", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
		"id":        data.ID.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read reads the workload resource.
func (r *WorkloadResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WorkloadResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading Workload", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Read workload from API
	var apiResp APIWorkload
	path := fmt.Sprintf("/config/namespaces/%s/workloads/%s", data.Namespace.ValueString(), data.Name.ValueString())
	err := r.client.Get(ctx, path, &apiResp)
	if err != nil {
		if client.IsNotFound(err) {
			tflog.Debug(ctx, "Workload not found, removing from state", map[string]interface{}{
				"name":      data.Name.ValueString(),
				"namespace": data.Namespace.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error Reading Workload",
			fmt.Sprintf("Could not read workload %s: %s", data.Name.ValueString(), err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update updates the workload resource.
func (r *WorkloadResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data WorkloadResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating Workload", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Build API request
	apiReq := data.ToAPIRequest()

	// Update workload via API
	var apiResp APIWorkload
	path := fmt.Sprintf("/config/namespaces/%s/workloads/%s", data.Namespace.ValueString(), data.Name.ValueString())
	err := r.client.Put(ctx, path, apiReq, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Workload",
			fmt.Sprintf("Could not update workload %s: %s", data.Name.ValueString(), err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete deletes the workload resource.
func (r *WorkloadResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data WorkloadResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting Workload", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	// Delete workload via API
	path := fmt.Sprintf("/config/namespaces/%s/workloads/%s", data.Namespace.ValueString(), data.Name.ValueString())
	err := r.client.Delete(ctx, path)
	if err != nil {
		if client.IsNotFound(err) {
			// Already deleted, nothing to do
			tflog.Debug(ctx, "Workload already deleted", map[string]interface{}{
				"name":      data.Name.ValueString(),
				"namespace": data.Namespace.ValueString(),
			})
			return
		}
		resp.Diagnostics.AddError(
			"Error Deleting Workload",
			fmt.Sprintf("Could not delete workload %s: %s", data.Name.ValueString(), err),
		)
		return
	}

	tflog.Debug(ctx, "Deleted Workload", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})
}

// ImportState imports an existing workload into Terraform state.
func (r *WorkloadResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Importing Workload", map[string]interface{}{
		"id": req.ID,
	})

	// Parse namespace/name from import ID
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("Expected namespace/name format, got: %s", req.ID),
		)
		return
	}

	namespace := parts[0]
	name := parts[1]

	// Read workload from API to verify it exists
	var apiResp APIWorkload
	path := fmt.Sprintf("/config/namespaces/%s/workloads/%s", namespace, name)
	err := r.client.Get(ctx, path, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Importing Workload",
			fmt.Sprintf("Could not import workload %s/%s: %s", namespace, name, err),
		)
		return
	}

	// Update model from response
	var data WorkloadResourceModel
	data.FromAPIResponse(&apiResp)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
