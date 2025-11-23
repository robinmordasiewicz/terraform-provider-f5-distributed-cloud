// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package namespace

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

// Default timeouts for namespace operations.
const (
	defaultCreateTimeout = 10 * time.Minute
	defaultReadTimeout   = 5 * time.Minute
	defaultUpdateTimeout = 10 * time.Minute
	defaultDeleteTimeout = 10 * time.Minute
)

// privateStateKey is the key for storing private state data.
const privateStateKey = "namespace_metadata"

// privateStateData stores API metadata for drift detection.
type privateStateData struct {
	UID               string `json:"uid,omitempty"`
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
	ModificationTime  string `json:"modification_timestamp,omitempty"`
}

// Ensure provider-defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                   = &NamespaceResource{}
	_ resource.ResourceWithImportState    = &NamespaceResource{}
	_ resource.ResourceWithModifyPlan     = &NamespaceResource{}
	_ resource.ResourceWithUpgradeState   = &NamespaceResource{}
	_ resource.ResourceWithValidateConfig = &NamespaceResource{}
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

	// Get create timeout
	createTimeout, diags := data.Timeouts.Create(ctx, defaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	tflog.Debug(ctx, "Creating namespace", map[string]interface{}{
		"name":    data.Name.ValueString(),
		"timeout": createTimeout.String(),
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

	// Save private state for drift detection
	r.savePrivateState(ctx, resp, &apiResp)

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

	// Get read timeout
	readTimeout, diags := data.Timeouts.Read(ctx, defaultReadTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, readTimeout)
	defer cancel()

	tflog.Debug(ctx, "Reading namespace", map[string]interface{}{
		"name":    data.Name.ValueString(),
		"timeout": readTimeout.String(),
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

	// Get update timeout
	updateTimeout, diags := data.Timeouts.Update(ctx, defaultUpdateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, updateTimeout)
	defer cancel()

	tflog.Debug(ctx, "Updating namespace", map[string]interface{}{
		"name":    data.Name.ValueString(),
		"timeout": updateTimeout.String(),
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

	// Get delete timeout
	deleteTimeout, diags := data.Timeouts.Delete(ctx, defaultDeleteTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, deleteTimeout)
	defer cancel()

	tflog.Debug(ctx, "Deleting namespace", map[string]interface{}{
		"name":    data.Name.ValueString(),
		"timeout": deleteTimeout.String(),
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

// ValidateConfig validates the resource configuration.
func (r *NamespaceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data NamespaceResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Additional validation can be added here if needed
	// For example, validating that certain field combinations are valid
}

// ModifyPlan modifies the plan to add warnings for unknown values.
func (r *NamespaceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip if destroying
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan NamespaceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Warn about unknown values during plan
	if plan.Name.IsUnknown() {
		resp.Diagnostics.AddWarning(
			"Unknown Namespace Name",
			"The namespace name contains an unknown value and will be determined during apply. "+
				"This may cause issues if other resources depend on this namespace.",
		)
	}
}

// UpgradeState handles state upgrades from older schema versions.
func (r *NamespaceResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// Version 0 to current version upgrade
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Required: true,
					},
					"description": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				// Read old state
				var oldState struct {
					ID          string `tfsdk:"id"`
					Name        string `tfsdk:"name"`
					Description string `tfsdk:"description"`
				}

				resp.Diagnostics.Append(req.State.Get(ctx, &oldState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				// Set new state with same values
				// New schema adds timeouts block which will be null by default
				resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), oldState.ID)...)
				resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), oldState.Name)...)
				resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("description"), oldState.Description)...)

				tflog.Info(ctx, "Upgraded namespace state from version 0", map[string]interface{}{
					"name": oldState.Name,
				})
			},
		},
	}
}

// savePrivateState saves API metadata to private state for drift detection.
func (r *NamespaceResource) savePrivateState(ctx context.Context, resp *resource.CreateResponse, apiResp *APINamespace) {
	privateData := privateStateData{
		UID:               apiResp.SystemMeta.UID,
		CreationTimestamp: apiResp.SystemMeta.CreationTimestamp,
		ModificationTime:  apiResp.SystemMeta.ModificationTime,
	}

	data, err := json.Marshal(privateData)
	if err != nil {
		tflog.Warn(ctx, "Failed to marshal private state data", map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	resp.Private.SetKey(ctx, privateStateKey, data)
}

// privateStateGetter is an interface for accessing private state from various request types.
type privateStateGetter interface {
	GetKey(ctx context.Context, key string) ([]byte, diag.Diagnostics)
}

// getPrivateState retrieves API metadata from private state.
func (r *NamespaceResource) getPrivateState(ctx context.Context, private privateStateGetter) (*privateStateData, error) {
	data, diags := private.GetKey(ctx, privateStateKey)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to get private state: %v", diags.Errors())
	}

	if data == nil {
		return nil, nil
	}

	var privateData privateStateData
	if err := json.Unmarshal(data, &privateData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal private state: %w", err)
	}

	return &privateData, nil
}
