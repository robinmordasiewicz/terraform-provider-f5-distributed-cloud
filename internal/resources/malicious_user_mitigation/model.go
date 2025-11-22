// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package malicious_user_mitigation

import "github.com/hashicorp/terraform-plugin-framework/types"

type MaliciousUserMitigationResourceModel struct {
	Name                types.String `tfsdk:"name"`
	Namespace           types.String `tfsdk:"namespace"`
	Description         types.String `tfsdk:"description"`
	MitigationType      types.String `tfsdk:"mitigation_type"`
	CaptchaChallenge    types.Bool   `tfsdk:"captcha_challenge"`
	JavascriptChallenge types.Bool   `tfsdk:"javascript_challenge"`
	TemporaryBlock      types.Bool   `tfsdk:"temporary_block"`
	BlockDuration       types.Int64  `tfsdk:"block_duration"`
	ID                  types.String `tfsdk:"id"`
}

type APIMaliciousUserMitigation struct {
	Metadata   APIMetadata                    `json:"metadata"`
	Spec       APIMaliciousUserMitigationSpec `json:"spec"`
	SystemMeta APISystemMetadata              `json:"system_metadata,omitempty"`
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

type APIMaliciousUserMitigationSpec struct {
	MitigationType      string `json:"mitigation_type,omitempty"`
	CaptchaChallenge    bool   `json:"captcha_challenge,omitempty"`
	JavascriptChallenge bool   `json:"javascript_challenge,omitempty"`
	TemporaryBlock      bool   `json:"temporary_block,omitempty"`
	BlockDuration       int64  `json:"block_duration,omitempty"`
}

func (m *MaliciousUserMitigationResourceModel) ToAPIRequest() *APIMaliciousUserMitigation {
	return &APIMaliciousUserMitigation{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIMaliciousUserMitigationSpec{
			MitigationType:      m.MitigationType.ValueString(),
			CaptchaChallenge:    m.CaptchaChallenge.ValueBool(),
			JavascriptChallenge: m.JavascriptChallenge.ValueBool(),
			TemporaryBlock:      m.TemporaryBlock.ValueBool(),
			BlockDuration:       m.BlockDuration.ValueInt64(),
		},
	}
}

func (m *MaliciousUserMitigationResourceModel) FromAPIResponse(resp *APIMaliciousUserMitigation) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	m.MitigationType = types.StringValue(resp.Spec.MitigationType)
	m.CaptchaChallenge = types.BoolValue(resp.Spec.CaptchaChallenge)
	m.JavascriptChallenge = types.BoolValue(resp.Spec.JavascriptChallenge)
	m.TemporaryBlock = types.BoolValue(resp.Spec.TemporaryBlock)
	m.BlockDuration = types.Int64Value(resp.Spec.BlockDuration)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
