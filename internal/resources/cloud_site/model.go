// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cloud_site

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// CloudSiteResourceModel describes the resource data model.
type CloudSiteResourceModel struct {
	ID                types.String      `tfsdk:"id"`
	Name              types.String      `tfsdk:"name"`
	Description       types.String      `tfsdk:"description"`
	SiteType          types.String      `tfsdk:"site_type"`
	Provider          types.String      `tfsdk:"cloud_provider"`
	Region            types.String      `tfsdk:"region"`
	VPCName           types.String      `tfsdk:"vpc_name"`
	VPCID             types.String      `tfsdk:"vpc_id"`
	CIDR              types.String      `tfsdk:"cidr"`
	Subnets           []SubnetModel     `tfsdk:"subnets"`
	MasterNodes       *MasterNodesModel `tfsdk:"master_nodes"`
	WorkerNodes       *WorkerNodesModel `tfsdk:"worker_nodes"`
	SSHKey            types.String      `tfsdk:"ssh_key"`
	Labels            types.Map         `tfsdk:"labels"`
	Annotations       types.Map         `tfsdk:"annotations"`
	SiteState         types.String      `tfsdk:"site_state"`
	SoftwareVersion   types.String      `tfsdk:"software_version"`
	OperatingSystem   types.String      `tfsdk:"operating_system"`
}

// SubnetModel describes a subnet configuration.
type SubnetModel struct {
	Name             types.String `tfsdk:"name"`
	CIDR             types.String `tfsdk:"cidr"`
	AvailabilityZone types.String `tfsdk:"availability_zone"`
	SubnetType       types.String `tfsdk:"subnet_type"`
}

// MasterNodesModel describes master node configuration.
type MasterNodesModel struct {
	InstanceType types.String `tfsdk:"instance_type"`
	DiskSize     types.Int64  `tfsdk:"disk_size"`
	Count        types.Int64  `tfsdk:"count"`
}

// WorkerNodesModel describes worker node configuration.
type WorkerNodesModel struct {
	InstanceType types.String `tfsdk:"instance_type"`
	DiskSize     types.Int64  `tfsdk:"disk_size"`
	Count        types.Int64  `tfsdk:"count"`
	MinCount     types.Int64  `tfsdk:"min_count"`
	MaxCount     types.Int64  `tfsdk:"max_count"`
}

// APICloudSite represents the API request/response structure.
type APICloudSite struct {
	Metadata    APIMetadata          `json:"metadata"`
	Spec        APICloudSiteSpec     `json:"spec"`
	SystemMeta  *APISystemMetadata   `json:"system_metadata,omitempty"`
	Status      *APICloudSiteStatus  `json:"status,omitempty"`
}

// APIMetadata represents the metadata section.
type APIMetadata struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// APISystemMetadata represents system-generated metadata.
type APISystemMetadata struct {
	UID string `json:"uid,omitempty"`
}

// APICloudSiteSpec represents the cloud site specification.
type APICloudSiteSpec struct {
	SiteType        string            `json:"site_type,omitempty"`
	CloudProvider   string            `json:"cloud_provider,omitempty"`
	Region          string            `json:"region,omitempty"`
	VPCName         string            `json:"vpc_name,omitempty"`
	VPCID           string            `json:"vpc_id,omitempty"`
	CIDR            string            `json:"cidr,omitempty"`
	Subnets         []APISubnet       `json:"subnets,omitempty"`
	MasterNodes     *APIMasterNodes   `json:"master_nodes,omitempty"`
	WorkerNodes     *APIWorkerNodes   `json:"worker_nodes,omitempty"`
	SSHKey          string            `json:"ssh_key,omitempty"`
	SoftwareVersion string            `json:"software_version,omitempty"`
	OperatingSystem string            `json:"operating_system,omitempty"`
}

// APISubnet represents subnet configuration in API.
type APISubnet struct {
	Name             string `json:"name,omitempty"`
	CIDR             string `json:"cidr,omitempty"`
	AvailabilityZone string `json:"availability_zone,omitempty"`
	SubnetType       string `json:"subnet_type,omitempty"`
}

// APIMasterNodes represents master node configuration in API.
type APIMasterNodes struct {
	InstanceType string `json:"instance_type,omitempty"`
	DiskSize     int64  `json:"disk_size,omitempty"`
	Count        int64  `json:"count,omitempty"`
}

// APIWorkerNodes represents worker node configuration in API.
type APIWorkerNodes struct {
	InstanceType string `json:"instance_type,omitempty"`
	DiskSize     int64  `json:"disk_size,omitempty"`
	Count        int64  `json:"count,omitempty"`
	MinCount     int64  `json:"min_count,omitempty"`
	MaxCount     int64  `json:"max_count,omitempty"`
}

// APICloudSiteStatus represents the status of a cloud site.
type APICloudSiteStatus struct {
	State string `json:"state,omitempty"`
}

// ToAPIRequest converts the Terraform model to API request format.
func (m *CloudSiteResourceModel) ToAPIRequest() *APICloudSite {
	apiReq := &APICloudSite{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APICloudSiteSpec{
			SiteType:        m.SiteType.ValueString(),
			CloudProvider:   m.Provider.ValueString(),
			Region:          m.Region.ValueString(),
			VPCName:         m.VPCName.ValueString(),
			VPCID:           m.VPCID.ValueString(),
			CIDR:            m.CIDR.ValueString(),
			SSHKey:          m.SSHKey.ValueString(),
			SoftwareVersion: m.SoftwareVersion.ValueString(),
			OperatingSystem: m.OperatingSystem.ValueString(),
		},
	}

	// Convert labels
	if !m.Labels.IsNull() {
		apiReq.Metadata.Labels = make(map[string]string)
		for k, v := range m.Labels.Elements() {
			if sv, ok := v.(types.String); ok {
				apiReq.Metadata.Labels[k] = sv.ValueString()
			}
		}
	}

	// Convert annotations
	if !m.Annotations.IsNull() {
		apiReq.Metadata.Annotations = make(map[string]string)
		for k, v := range m.Annotations.Elements() {
			if sv, ok := v.(types.String); ok {
				apiReq.Metadata.Annotations[k] = sv.ValueString()
			}
		}
	}

	// Convert subnets
	if len(m.Subnets) > 0 {
		apiReq.Spec.Subnets = make([]APISubnet, len(m.Subnets))
		for i, subnet := range m.Subnets {
			apiReq.Spec.Subnets[i] = APISubnet{
				Name:             subnet.Name.ValueString(),
				CIDR:             subnet.CIDR.ValueString(),
				AvailabilityZone: subnet.AvailabilityZone.ValueString(),
				SubnetType:       subnet.SubnetType.ValueString(),
			}
		}
	}

	// Convert master nodes
	if m.MasterNodes != nil {
		apiReq.Spec.MasterNodes = &APIMasterNodes{
			InstanceType: m.MasterNodes.InstanceType.ValueString(),
			DiskSize:     m.MasterNodes.DiskSize.ValueInt64(),
			Count:        m.MasterNodes.Count.ValueInt64(),
		}
	}

	// Convert worker nodes
	if m.WorkerNodes != nil {
		apiReq.Spec.WorkerNodes = &APIWorkerNodes{
			InstanceType: m.WorkerNodes.InstanceType.ValueString(),
			DiskSize:     m.WorkerNodes.DiskSize.ValueInt64(),
			Count:        m.WorkerNodes.Count.ValueInt64(),
			MinCount:     m.WorkerNodes.MinCount.ValueInt64(),
			MaxCount:     m.WorkerNodes.MaxCount.ValueInt64(),
		}
	}

	return apiReq
}

// FromAPIResponse updates the Terraform model from API response.
func (m *CloudSiteResourceModel) FromAPIResponse(resp *APICloudSite) {
	m.Name = types.StringValue(resp.Metadata.Name)

	if resp.Metadata.Description != "" {
		m.Description = types.StringValue(resp.Metadata.Description)
	}

	if resp.SystemMeta != nil && resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Name)
	}

	// Update spec fields
	if resp.Spec.SiteType != "" {
		m.SiteType = types.StringValue(resp.Spec.SiteType)
	}
	if resp.Spec.CloudProvider != "" {
		m.Provider = types.StringValue(resp.Spec.CloudProvider)
	}
	if resp.Spec.Region != "" {
		m.Region = types.StringValue(resp.Spec.Region)
	}
	if resp.Spec.VPCName != "" {
		m.VPCName = types.StringValue(resp.Spec.VPCName)
	}
	if resp.Spec.VPCID != "" {
		m.VPCID = types.StringValue(resp.Spec.VPCID)
	}
	if resp.Spec.CIDR != "" {
		m.CIDR = types.StringValue(resp.Spec.CIDR)
	}
	if resp.Spec.SSHKey != "" {
		m.SSHKey = types.StringValue(resp.Spec.SSHKey)
	}
	if resp.Spec.SoftwareVersion != "" {
		m.SoftwareVersion = types.StringValue(resp.Spec.SoftwareVersion)
	}
	if resp.Spec.OperatingSystem != "" {
		m.OperatingSystem = types.StringValue(resp.Spec.OperatingSystem)
	}

	// Update status
	if resp.Status != nil && resp.Status.State != "" {
		m.SiteState = types.StringValue(resp.Status.State)
	}
}
