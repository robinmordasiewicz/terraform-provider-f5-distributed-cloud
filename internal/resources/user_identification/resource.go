// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package user_identification

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/client"
)

var (
	_ resource.Resource                = &UserIdentificationResource{}
	_ resource.ResourceWithImportState = &UserIdentificationResource{}
)

// NewUserIdentificationResource creates a new User Identification resource.
func NewUserIdentificationResource() resource.Resource {
	return &UserIdentificationResource{}
}

// UserIdentificationResource defines the resource implementation.
type UserIdentificationResource struct {
	client *client.Client
}

// Metadata returns the resource type name.
func (r *UserIdentificationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_identification"
}

// Schema returns the resource schema.
func (r *UserIdentificationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

// Configure configures the resource with the provider client.
func (r *UserIdentificationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create creates a new User Identification resource.
func (r *UserIdentificationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data UserIdentificationResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating User Identification", map[string]interface{}{
		"name":      data.Name.ValueString(),
		"namespace": data.Namespace.ValueString(),
	})

	apiReq := data.ToAPIRequest()

	var apiResp APIUserIdentification
	path := fmt.Sprintf("/config/namespaces/%s/user_identifications", data.Namespace.ValueString())
	if err := r.client.Post(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating User Identification",
			fmt.Sprintf("Could not create User Identification %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err),
		)
		return
	}

	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read reads the User Identification resource.
func (r *UserIdentificationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data UserIdentificationResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APIUserIdentification
	path := fmt.Sprintf("/config/namespaces/%s/user_identifications/%s",
		data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error Reading User Identification",
			fmt.Sprintf("Could not read User Identification %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err),
		)
		return
	}

	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update updates the User Identification resource.
func (r *UserIdentificationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data UserIdentificationResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := data.ToAPIRequest()

	var apiResp APIUserIdentification
	path := fmt.Sprintf("/config/namespaces/%s/user_identifications/%s",
		data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Put(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating User Identification",
			fmt.Sprintf("Could not update User Identification %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err),
		)
		return
	}

	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete deletes the User Identification resource.
func (r *UserIdentificationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data UserIdentificationResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := fmt.Sprintf("/config/namespaces/%s/user_identifications/%s",
		data.Namespace.ValueString(), data.Name.ValueString())
	if err := r.client.Delete(ctx, path); err != nil {
		if client.IsNotFound(err) {
			return
		}
		resp.Diagnostics.AddError(
			"Error Deleting User Identification",
			fmt.Sprintf("Could not delete User Identification %s/%s: %s",
				data.Namespace.ValueString(), data.Name.ValueString(), err),
		)
		return
	}
}

// ImportState imports an existing User Identification resource.
func (r *UserIdentificationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Invalid Import ID",
			fmt.Sprintf("Expected format 'namespace/name', got: %s", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("namespace"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), parts[1])...)
}
