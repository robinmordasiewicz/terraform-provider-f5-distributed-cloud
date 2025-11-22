// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package discovery

import "github.com/hashicorp/terraform-plugin-framework/types"

type DiscoveryResourceModel struct {
	Name            types.String `tfsdk:"name"`
	Namespace       types.String `tfsdk:"namespace"`
	Description     types.String `tfsdk:"description"`
	DiscoveryType   types.String `tfsdk:"discovery_type"`
	ClusterName     types.String `tfsdk:"cluster_name"`
	KubeConfig      types.String `tfsdk:"kube_config"`
	VirtualSiteRef  types.String `tfsdk:"virtual_site_ref"`
	PublishInfo     *PublishInfoModel `tfsdk:"publish_info"`
	ID              types.String `tfsdk:"id"`
}

type PublishInfoModel struct {
	DisablePublish types.Bool   `tfsdk:"disable_publish"`
	DNSDomain      types.String `tfsdk:"dns_domain"`
}

type APIDiscovery struct {
	Metadata   APIMetadata       `json:"metadata"`
	Spec       APIDiscoverySpec  `json:"spec"`
	SystemMeta APISystemMetadata `json:"system_metadata,omitempty"`
}

type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace,omitempty"`
	Description string `json:"description,omitempty"`
}

type APISystemMetadata struct {
	UID               string `json:"uid,omitempty"`
	CreationTimestamp string `json:"creation_timestamp,omitempty"`
}

type APIDiscoverySpec struct {
	DiscoveryType  string          `json:"discovery_type,omitempty"`
	ClusterName    string          `json:"cluster_name,omitempty"`
	KubeConfig     string          `json:"kube_config,omitempty"`
	VirtualSiteRef string          `json:"virtual_site_ref,omitempty"`
	PublishInfo    *APIPublishInfo `json:"publish_info,omitempty"`
}

type APIPublishInfo struct {
	DisablePublish bool   `json:"disable_publish,omitempty"`
	DNSDomain      string `json:"dns_domain,omitempty"`
}

func (m *DiscoveryResourceModel) ToAPIRequest() *APIDiscovery {
	spec := APIDiscoverySpec{}
	if !m.DiscoveryType.IsNull() {
		spec.DiscoveryType = m.DiscoveryType.ValueString()
	}
	if !m.ClusterName.IsNull() {
		spec.ClusterName = m.ClusterName.ValueString()
	}
	if !m.KubeConfig.IsNull() {
		spec.KubeConfig = m.KubeConfig.ValueString()
	}
	if !m.VirtualSiteRef.IsNull() {
		spec.VirtualSiteRef = m.VirtualSiteRef.ValueString()
	}
	if m.PublishInfo != nil {
		spec.PublishInfo = &APIPublishInfo{
			DisablePublish: m.PublishInfo.DisablePublish.ValueBool(),
			DNSDomain:      m.PublishInfo.DNSDomain.ValueString(),
		}
	}
	return &APIDiscovery{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *DiscoveryResourceModel) FromAPIResponse(resp *APIDiscovery) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.Spec.DiscoveryType != "" {
		m.DiscoveryType = types.StringValue(resp.Spec.DiscoveryType)
	}
	if resp.Spec.ClusterName != "" {
		m.ClusterName = types.StringValue(resp.Spec.ClusterName)
	}
	if resp.Spec.VirtualSiteRef != "" {
		m.VirtualSiteRef = types.StringValue(resp.Spec.VirtualSiteRef)
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
