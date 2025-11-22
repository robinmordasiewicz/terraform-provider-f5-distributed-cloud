// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cloud_credentials

import "github.com/hashicorp/terraform-plugin-framework/types"

type CloudCredentialsResourceModel struct {
	Name             types.String           `tfsdk:"name"`
	Namespace        types.String           `tfsdk:"namespace"`
	Description      types.String           `tfsdk:"description"`
	CloudProvider    types.String           `tfsdk:"cloud_provider"`
	AWSCredentials   *AWSCredentialsModel   `tfsdk:"aws_credentials"`
	AzureCredentials *AzureCredentialsModel `tfsdk:"azure_credentials"`
	GCPCredentials   *GCPCredentialsModel   `tfsdk:"gcp_credentials"`
	ID               types.String           `tfsdk:"id"`
}

type AWSCredentialsModel struct {
	AccessKeyID     types.String `tfsdk:"access_key_id"`
	SecretAccessKey types.String `tfsdk:"secret_access_key"`
}

type AzureCredentialsModel struct {
	SubscriptionID types.String `tfsdk:"subscription_id"`
	TenantID       types.String `tfsdk:"tenant_id"`
	ClientID       types.String `tfsdk:"client_id"`
	ClientSecret   types.String `tfsdk:"client_secret"`
}

type GCPCredentialsModel struct {
	ProjectID       types.String `tfsdk:"project_id"`
	CredentialsFile types.String `tfsdk:"credentials_file"`
}

type APICloudCredentials struct {
	Metadata   APIMetadata             `json:"metadata"`
	Spec       APICloudCredentialsSpec `json:"spec"`
	SystemMeta APISystemMetadata       `json:"system_metadata,omitempty"`
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

type APICloudCredentialsSpec struct {
	CloudProvider string        `json:"cloud_provider,omitempty"`
	AWSCred       *APIAWSCred   `json:"aws_cred,omitempty"`
	AzureCred     *APIAzureCred `json:"azure_cred,omitempty"`
	GCPCred       *APIGCPCred   `json:"gcp_cred,omitempty"`
}

type APIAWSCred struct {
	AccessKey string `json:"access_key,omitempty"`
	SecretKey string `json:"secret_key,omitempty"`
}

type APIAzureCred struct {
	SubscriptionID string `json:"subscription_id,omitempty"`
	TenantID       string `json:"tenant_id,omitempty"`
	ClientID       string `json:"client_id,omitempty"`
	ClientSecret   string `json:"client_secret,omitempty"`
}

type APIGCPCred struct {
	ProjectID       string `json:"project_id,omitempty"`
	CredentialsFile string `json:"credentials_file,omitempty"`
}

func (m *CloudCredentialsResourceModel) ToAPIRequest() *APICloudCredentials {
	spec := APICloudCredentialsSpec{}
	if !m.CloudProvider.IsNull() {
		spec.CloudProvider = m.CloudProvider.ValueString()
	}
	if m.AWSCredentials != nil {
		spec.AWSCred = &APIAWSCred{
			AccessKey: m.AWSCredentials.AccessKeyID.ValueString(),
			SecretKey: m.AWSCredentials.SecretAccessKey.ValueString(),
		}
	}
	if m.AzureCredentials != nil {
		spec.AzureCred = &APIAzureCred{
			SubscriptionID: m.AzureCredentials.SubscriptionID.ValueString(),
			TenantID:       m.AzureCredentials.TenantID.ValueString(),
			ClientID:       m.AzureCredentials.ClientID.ValueString(),
			ClientSecret:   m.AzureCredentials.ClientSecret.ValueString(),
		}
	}
	if m.GCPCredentials != nil {
		spec.GCPCred = &APIGCPCred{
			ProjectID:       m.GCPCredentials.ProjectID.ValueString(),
			CredentialsFile: m.GCPCredentials.CredentialsFile.ValueString(),
		}
	}
	return &APICloudCredentials{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *CloudCredentialsResourceModel) FromAPIResponse(resp *APICloudCredentials) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.Spec.CloudProvider != "" {
		m.CloudProvider = types.StringValue(resp.Spec.CloudProvider)
	}
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
