package waf_exclusion_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// WAFExclusionPolicyResourceModel describes the resource data model.
type WAFExclusionPolicyResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Namespace   types.String `tfsdk:"namespace"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
}

// APIMetadata represents the metadata structure in API requests/responses.
type APIMetadata struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Description string `json:"description,omitempty"`
}

// APIWAFExclusionPolicySpec represents the spec structure in API requests/responses.
type APIWAFExclusionPolicySpec struct{}

// APISystemMetadata represents system-generated metadata in API responses.
type APISystemMetadata struct {
	UID string `json:"uid,omitempty"`
}

// APIWAFExclusionPolicy represents the full API request/response structure.
type APIWAFExclusionPolicy struct {
	Metadata       APIMetadata               `json:"metadata"`
	Spec           APIWAFExclusionPolicySpec `json:"spec"`
	SystemMetadata APISystemMetadata         `json:"system_metadata,omitempty"`
}

// ToAPIRequest converts the Terraform model to an API request structure.
func (m *WAFExclusionPolicyResourceModel) ToAPIRequest() APIWAFExclusionPolicy {
	return APIWAFExclusionPolicy{
		Metadata: APIMetadata{
			Name:        m.Name.ValueString(),
			Namespace:   m.Namespace.ValueString(),
			Description: m.Description.ValueString(),
		},
		Spec: APIWAFExclusionPolicySpec{},
	}
}

// FromAPIResponse updates the Terraform model from an API response.
func (m *WAFExclusionPolicyResourceModel) FromAPIResponse(api APIWAFExclusionPolicy) {
	m.Name = types.StringValue(api.Metadata.Name)
	m.Namespace = types.StringValue(api.Metadata.Namespace)
	m.Description = types.StringValue(api.Metadata.Description)
	if api.SystemMetadata.UID != "" {
		m.ID = types.StringValue(api.SystemMetadata.UID)
	}
}
