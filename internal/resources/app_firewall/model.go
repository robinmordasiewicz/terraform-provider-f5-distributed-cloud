// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package app_firewall

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// AppFirewallResourceModel describes the resource data model.
type AppFirewallResourceModel struct {
	ID                     types.String             `tfsdk:"id"`
	Name                   types.String             `tfsdk:"name"`
	Namespace              types.String             `tfsdk:"namespace"`
	Description            types.String             `tfsdk:"description"`
	Mode                   types.String             `tfsdk:"mode"`
	DetectionSettings      *DetectionSettingsModel  `tfsdk:"detection_settings"`
	BotProtection          *BotProtectionModel      `tfsdk:"bot_protection"`
	AllowedResponseCodes   types.List               `tfsdk:"allowed_response_codes"`
	DefaultAnonymization   types.Bool               `tfsdk:"default_anonymization"`
	UseDefaultBlockingPage types.Bool               `tfsdk:"use_default_blocking_page"`
	CustomBlockingPage     *CustomBlockingPageModel `tfsdk:"custom_blocking_page"`
	Labels                 types.Map                `tfsdk:"labels"`
}

// DetectionSettingsModel describes detection settings.
type DetectionSettingsModel struct {
	SignatureBasedBotProtection types.Bool   `tfsdk:"signature_based_bot_protection"`
	XhrCheck                    types.Bool   `tfsdk:"xhr_check"`
	CookieProtection            types.Bool   `tfsdk:"cookie_protection"`
	EnableSuppression           types.Bool   `tfsdk:"enable_suppression"`
	ViolationRating             types.String `tfsdk:"violation_rating"`
}

// BotProtectionModel describes bot protection settings.
type BotProtectionModel struct {
	GoodBotAction       types.String `tfsdk:"good_bot_action"`
	MaliciousBotAction  types.String `tfsdk:"malicious_bot_action"`
	SuspiciousBotAction types.String `tfsdk:"suspicious_bot_action"`
}

// CustomBlockingPageModel describes custom blocking page settings.
type CustomBlockingPageModel struct {
	BlockingPageBody types.String `tfsdk:"blocking_page_body"`
	ResponseCode     types.String `tfsdk:"response_code"`
}

// APIAppFirewall represents the API request/response structure.
type APIAppFirewall struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APIAppFirewallSpec `json:"spec"`
	SystemMeta *APISystemMetadata `json:"system_metadata,omitempty"`
}

// APIMetadata represents the metadata section.
type APIMetadata struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace,omitempty"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

// APISystemMetadata represents system-generated metadata.
type APISystemMetadata struct {
	UID string `json:"uid,omitempty"`
}

// APIAppFirewallSpec represents the app firewall specification.
type APIAppFirewallSpec struct {
	Mode                   string                 `json:"mode,omitempty"`
	DetectionSettings      *APIDetectionSettings  `json:"detection_settings,omitempty"`
	BotProtection          *APIBotProtection      `json:"bot_protection,omitempty"`
	AllowedResponseCodes   []string               `json:"allowed_response_codes,omitempty"`
	DefaultAnonymization   bool                   `json:"default_anonymization,omitempty"`
	UseDefaultBlockingPage bool                   `json:"use_default_blocking_page,omitempty"`
	CustomBlockingPage     *APICustomBlockingPage `json:"custom_blocking_page,omitempty"`
}

// APIDetectionSettings represents detection settings in API.
type APIDetectionSettings struct {
	SignatureBasedBotProtection bool   `json:"signature_based_bot_protection,omitempty"`
	XhrCheck                    bool   `json:"xhr_check,omitempty"`
	CookieProtection            bool   `json:"cookie_protection,omitempty"`
	EnableSuppression           bool   `json:"enable_suppression,omitempty"`
	ViolationRating             string `json:"violation_rating,omitempty"`
}

// APIBotProtection represents bot protection settings in API.
type APIBotProtection struct {
	GoodBotAction       string `json:"good_bot_action,omitempty"`
	MaliciousBotAction  string `json:"malicious_bot_action,omitempty"`
	SuspiciousBotAction string `json:"suspicious_bot_action,omitempty"`
}

// APICustomBlockingPage represents custom blocking page in API.
type APICustomBlockingPage struct {
	BlockingPageBody string `json:"blocking_page_body,omitempty"`
	ResponseCode     string `json:"response_code,omitempty"`
}

// ToAPIRequest converts the Terraform model to API request format.
func (m *AppFirewallResourceModel) ToAPIRequest() *APIAppFirewall {
	apiReq := &APIAppFirewall{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIAppFirewallSpec{
			Mode:                   m.Mode.ValueString(),
			DefaultAnonymization:   m.DefaultAnonymization.ValueBool(),
			UseDefaultBlockingPage: m.UseDefaultBlockingPage.ValueBool(),
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

	// Convert allowed response codes
	if !m.AllowedResponseCodes.IsNull() {
		codes := make([]string, 0, len(m.AllowedResponseCodes.Elements()))
		for _, v := range m.AllowedResponseCodes.Elements() {
			if sv, ok := v.(types.String); ok {
				codes = append(codes, sv.ValueString())
			}
		}
		apiReq.Spec.AllowedResponseCodes = codes
	}

	// Convert detection settings
	if m.DetectionSettings != nil {
		apiReq.Spec.DetectionSettings = &APIDetectionSettings{
			SignatureBasedBotProtection: m.DetectionSettings.SignatureBasedBotProtection.ValueBool(),
			XhrCheck:                    m.DetectionSettings.XhrCheck.ValueBool(),
			CookieProtection:            m.DetectionSettings.CookieProtection.ValueBool(),
			EnableSuppression:           m.DetectionSettings.EnableSuppression.ValueBool(),
			ViolationRating:             m.DetectionSettings.ViolationRating.ValueString(),
		}
	}

	// Convert bot protection
	if m.BotProtection != nil {
		apiReq.Spec.BotProtection = &APIBotProtection{
			GoodBotAction:       m.BotProtection.GoodBotAction.ValueString(),
			MaliciousBotAction:  m.BotProtection.MaliciousBotAction.ValueString(),
			SuspiciousBotAction: m.BotProtection.SuspiciousBotAction.ValueString(),
		}
	}

	// Convert custom blocking page
	if m.CustomBlockingPage != nil {
		apiReq.Spec.CustomBlockingPage = &APICustomBlockingPage{
			BlockingPageBody: m.CustomBlockingPage.BlockingPageBody.ValueString(),
			ResponseCode:     m.CustomBlockingPage.ResponseCode.ValueString(),
		}
	}

	return apiReq
}

// FromAPIResponse updates the Terraform model from API response.
func (m *AppFirewallResourceModel) FromAPIResponse(resp *APIAppFirewall) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)

	if resp.Metadata.Description != "" {
		m.Description = types.StringValue(resp.Metadata.Description)
	}

	if resp.SystemMeta != nil && resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}

	if resp.Spec.Mode != "" {
		m.Mode = types.StringValue(resp.Spec.Mode)
	}

	m.DefaultAnonymization = types.BoolValue(resp.Spec.DefaultAnonymization)
	m.UseDefaultBlockingPage = types.BoolValue(resp.Spec.UseDefaultBlockingPage)
}
