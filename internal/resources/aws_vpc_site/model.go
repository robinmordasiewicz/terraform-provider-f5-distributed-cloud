// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package aws_vpc_site

import "github.com/hashicorp/terraform-plugin-framework/types"

type AWSVPCSiteResourceModel struct {
	Name                types.String `tfsdk:"name"`
	Namespace           types.String `tfsdk:"namespace"`
	Description         types.String `tfsdk:"description"`
	AWSRegion           types.String `tfsdk:"aws_region"`
	VPCID               types.String `tfsdk:"vpc_id"`
	InstanceType        types.String `tfsdk:"instance_type"`
	CloudCredentialsRef types.String `tfsdk:"cloud_credentials_ref"`
	ID                  types.String `tfsdk:"id"`
}

type APIAWSVPCSite struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APIAWSVPCSiteSpec  `json:"spec"`
	SystemMeta APISystemMetadata  `json:"system_metadata,omitempty"`
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

type APIAWSVPCSiteSpec struct {
	AWSRegion           string `json:"aws_region,omitempty"`
	VPCID               string `json:"vpc_id,omitempty"`
	InstanceType        string `json:"instance_type,omitempty"`
	CloudCredentialsRef string `json:"cloud_credentials_ref,omitempty"`
}

func (m *AWSVPCSiteResourceModel) ToAPIRequest() *APIAWSVPCSite {
	return &APIAWSVPCSite{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIAWSVPCSiteSpec{
			AWSRegion:           m.AWSRegion.ValueString(),
			VPCID:               m.VPCID.ValueString(),
			InstanceType:        m.InstanceType.ValueString(),
			CloudCredentialsRef: m.CloudCredentialsRef.ValueString(),
		},
	}
}

func (m *AWSVPCSiteResourceModel) FromAPIResponse(resp *APIAWSVPCSite) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.AWSRegion = types.StringValue(resp.Spec.AWSRegion)
	m.VPCID = types.StringValue(resp.Spec.VPCID)
	m.InstanceType = types.StringValue(resp.Spec.InstanceType)
	m.CloudCredentialsRef = types.StringValue(resp.Spec.CloudCredentialsRef)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
