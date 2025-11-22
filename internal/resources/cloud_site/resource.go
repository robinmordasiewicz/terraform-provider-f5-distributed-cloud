// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cloud_site

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
)

var (
	_ resource.Resource                = &CloudSiteResource{}
	_ resource.ResourceWithImportState = &CloudSiteResource{}
)

func NewCloudSiteResource() resource.Resource {
	return &CloudSiteResource{}
}

type CloudSiteResource struct {
	client *client.Client
}

func (r *CloudSiteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_site"
}

func (r *CloudSiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = Schema()
}

func (r *CloudSiteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// getSiteAPIPath returns the appropriate API path based on site type.
func getSiteAPIPath(siteType, name string) string {
	switch siteType {
	case "aws_vpc_site":
		if name == "" {
			return "/config/namespaces/system/aws_vpc_sites"
		}
		return fmt.Sprintf("/config/namespaces/system/aws_vpc_sites/%s", name)
	case "azure_vnet_site":
		if name == "" {
			return "/config/namespaces/system/azure_vnet_sites"
		}
		return fmt.Sprintf("/config/namespaces/system/azure_vnet_sites/%s", name)
	case "gcp_vpc_site":
		if name == "" {
			return "/config/namespaces/system/gcp_vpc_sites"
		}
		return fmt.Sprintf("/config/namespaces/system/gcp_vpc_sites/%s", name)
	default:
		if name == "" {
			return "/config/namespaces/system/sites"
		}
		return fmt.Sprintf("/config/namespaces/system/sites/%s", name)
	}
}

func (r *CloudSiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudSiteResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating Cloud Site", map[string]interface{}{
		"name": data.Name.ValueString(), "site_type": data.SiteType.ValueString(),
	})

	apiReq := data.ToAPIRequest()
	var apiResp APICloudSite
	path := getSiteAPIPath(data.SiteType.ValueString(), "")
	if err := r.client.Post(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Creating Cloud Site",
			fmt.Sprintf("Could not create cloud site %s: %s", data.Name.ValueString(), err))
		return
	}

	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudSiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudSiteResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var apiResp APICloudSite
	path := getSiteAPIPath(data.SiteType.ValueString(), data.Name.ValueString())
	if err := r.client.Get(ctx, path, &apiResp); err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading Cloud Site",
			fmt.Sprintf("Could not read cloud site %s: %s", data.Name.ValueString(), err))
		return
	}

	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudSiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CloudSiteResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := data.ToAPIRequest()
	var apiResp APICloudSite
	path := getSiteAPIPath(data.SiteType.ValueString(), data.Name.ValueString())
	if err := r.client.Put(ctx, path, apiReq, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error Updating Cloud Site",
			fmt.Sprintf("Could not update cloud site %s: %s", data.Name.ValueString(), err))
		return
	}

	data.FromAPIResponse(&apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudSiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudSiteResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := getSiteAPIPath(data.SiteType.ValueString(), data.Name.ValueString())
	if err := r.client.Delete(ctx, path); err != nil {
		if client.IsNotFound(err) {
			return
		}
		resp.Diagnostics.AddError("Error Deleting Cloud Site",
			fmt.Sprintf("Could not delete cloud site %s: %s", data.Name.ValueString(), err))
	}
}

func (r *CloudSiteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: name (sites are in system namespace)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), req.ID)...)
	// Note: site_type must be set after import via configuration
}
