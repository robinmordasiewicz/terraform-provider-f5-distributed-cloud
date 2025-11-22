// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package namespace

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

// Ensure provider-defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &NamespaceResource{}
	_ resource.ResourceWithImportState = &NamespaceResource{}
)

// NewNamespaceResource creates a new namespace resource.
func NewNamespaceResource() resource.Resource {
	return &NamespaceResource{}
}

// NamespaceResource defines the resource implementation.
type NamespaceResource struct {
	client *client.Client
}

// Metadata returns the resource type name.
func (r *NamespaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_namespace"
}

// Schema returns the resource schema.
func (r *NamespaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

// Configure adds the provider configured client to the resource.
func (r *NamespaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create creates the namespace resource.
func (r *NamespaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NamespaceResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating namespace", map[string]interface{}{
		"name": data.Name.ValueString(),
	})

	// Build API request
	apiReq := data.ToAPIRequest()

	// Create namespace via API
	var apiResp APINamespace
	err := r.client.Post(ctx, "/web/namespaces", apiReq, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Namespace",
			fmt.Sprintf("Could not create namespace %s: %s", data.Name.ValueString(), err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	tflog.Debug(ctx, "Created namespace", map[string]interface{}{
		"name": data.Name.ValueString(),
		"id":   data.ID.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read reads the namespace resource.
func (r *NamespaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NamespaceResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading namespace", map[string]interface{}{
		"name": data.Name.ValueString(),
	})

	// Read namespace from API
	var apiResp APINamespace
	err := r.client.Get(ctx, fmt.Sprintf("/web/namespaces/%s", data.Name.ValueString()), &apiResp)
	if err != nil {
		if client.IsNotFound(err) {
			tflog.Debug(ctx, "Namespace not found, removing from state", map[string]interface{}{
				"name": data.Name.ValueString(),
			})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error Reading Namespace",
			fmt.Sprintf("Could not read namespace %s: %s", data.Name.ValueString(), err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update updates the namespace resource.
func (r *NamespaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NamespaceResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating namespace", map[string]interface{}{
		"name": data.Name.ValueString(),
	})

	// Build API request
	apiReq := data.ToAPIRequest()

	// Update namespace via API
	var apiResp APINamespace
	err := r.client.Put(ctx, fmt.Sprintf("/web/namespaces/%s", data.Name.ValueString()), apiReq, &apiResp)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Namespace",
			fmt.Sprintf("Could not update namespace %s: %s", data.Name.ValueString(), err),
		)
		return
	}

	// Update model from response
	data.FromAPIResponse(&apiResp)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete deletes the namespace resource.
func (r *NamespaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NamespaceResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting namespace", map[string]interface{}{
		"name": data.Name.ValueString(),
	})

	// Delete namespace via API
	err := r.client.Delete(ctx, fmt.Sprintf("/web/namespaces/%s", data.Name.ValueString()))
	if err != nil {
		if client.IsNotFound(err) {
			// Already deleted, nothing to do
			tflog.Debug(ctx, "Namespace already deleted", map[string]interface{}{
				"name": data.Name.ValueString(),
			})
			return
		}
		resp.Diagnostics.AddError(
			"Error Deleting Namespace",
			fmt.Sprintf("Could not delete namespace %s: %s", data.Name.ValueString(), err),
		)
		return
	}

	tflog.Debug(ctx, "Deleted namespace", map[string]interface{}{
		"name": data.Name.ValueString(),
	})
}

// ImportState imports an existing namespace into Terraform state.
func (r *NamespaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Importing namespace", map[string]interface{}{
		"id": req.ID,
	})

	// Use the import ID as the name
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), req.ID)...)
}
