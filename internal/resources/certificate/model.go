// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// CertificateResourceModel describes the resource data model.
type CertificateResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`

	// Certificate data
	CertificateURL types.String `tfsdk:"certificate_url"`
	PrivateKey     *SecretModel `tfsdk:"private_key"`

	// Computed fields
	ID         types.String `tfsdk:"id"`
	Expiration types.String `tfsdk:"expiration"`
	Subject    types.String `tfsdk:"subject"`
	Issuer     types.String `tfsdk:"issuer"`
}

// SecretModel represents a secret reference.
type SecretModel struct {
	ClearSecretURL     types.String `tfsdk:"clear_secret_url"`
	BlindfoldSecretURL types.String `tfsdk:"blindfold_secret_url"`
	VaultSecretURL     types.String `tfsdk:"vault_secret_url"`
	WingmanSecretURL   types.String `tfsdk:"wingman_secret_url"`
}

// API structures

type APICertificate struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APICertificateSpec `json:"spec"`
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

type APICertificateSpec struct {
	CertificateURL string     `json:"certificate_url,omitempty"`
	PrivateKey     *APISecret `json:"private_key,omitempty"`
}

type APISecret struct {
	ClearSecretURL     string `json:"clear_secret_url,omitempty"`
	BlindfoldSecretURL string `json:"blindfold_secret_url,omitempty"`
	VaultSecretURL     string `json:"vault_secret_url,omitempty"`
	WingmanSecretURL   string `json:"wingman_secret_url,omitempty"`
}

// ToAPIRequest converts the Terraform model to an API request.
func (m *CertificateResourceModel) ToAPIRequest() *APICertificate {
	spec := APICertificateSpec{}

	if !m.CertificateURL.IsNull() && !m.CertificateURL.IsUnknown() {
		spec.CertificateURL = m.CertificateURL.ValueString()
	}

	if m.PrivateKey != nil {
		spec.PrivateKey = &APISecret{}
		if !m.PrivateKey.ClearSecretURL.IsNull() && !m.PrivateKey.ClearSecretURL.IsUnknown() {
			spec.PrivateKey.ClearSecretURL = m.PrivateKey.ClearSecretURL.ValueString()
		}
		if !m.PrivateKey.BlindfoldSecretURL.IsNull() && !m.PrivateKey.BlindfoldSecretURL.IsUnknown() {
			spec.PrivateKey.BlindfoldSecretURL = m.PrivateKey.BlindfoldSecretURL.ValueString()
		}
		if !m.PrivateKey.VaultSecretURL.IsNull() && !m.PrivateKey.VaultSecretURL.IsUnknown() {
			spec.PrivateKey.VaultSecretURL = m.PrivateKey.VaultSecretURL.ValueString()
		}
		if !m.PrivateKey.WingmanSecretURL.IsNull() && !m.PrivateKey.WingmanSecretURL.IsUnknown() {
			spec.PrivateKey.WingmanSecretURL = m.PrivateKey.WingmanSecretURL.ValueString()
		}
	}

	return &APICertificate{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: spec,
	}
}

// FromAPIResponse updates the model from an API response.
func (m *CertificateResourceModel) FromAPIResponse(resp *APICertificate) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)

	if resp.Spec.CertificateURL != "" {
		m.CertificateURL = types.StringValue(resp.Spec.CertificateURL)
	}

	if resp.Spec.PrivateKey != nil {
		m.PrivateKey = &SecretModel{}
		if resp.Spec.PrivateKey.ClearSecretURL != "" {
			m.PrivateKey.ClearSecretURL = types.StringValue(resp.Spec.PrivateKey.ClearSecretURL)
		}
		if resp.Spec.PrivateKey.BlindfoldSecretURL != "" {
			m.PrivateKey.BlindfoldSecretURL = types.StringValue(resp.Spec.PrivateKey.BlindfoldSecretURL)
		}
		if resp.Spec.PrivateKey.VaultSecretURL != "" {
			m.PrivateKey.VaultSecretURL = types.StringValue(resp.Spec.PrivateKey.VaultSecretURL)
		}
		if resp.Spec.PrivateKey.WingmanSecretURL != "" {
			m.PrivateKey.WingmanSecretURL = types.StringValue(resp.Spec.PrivateKey.WingmanSecretURL)
		}
	}

	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
