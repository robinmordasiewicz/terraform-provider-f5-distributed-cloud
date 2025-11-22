// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package alert_policy

import "github.com/hashicorp/terraform-plugin-framework/types"

type AlertPolicyResourceModel struct {
	Name        types.String       `tfsdk:"name"`
	Namespace   types.String       `tfsdk:"namespace"`
	Description types.String       `tfsdk:"description"`
	Receivers   []ReceiverRefModel `tfsdk:"receivers"`
	Routes      []RouteModel       `tfsdk:"routes"`
	ID          types.String       `tfsdk:"id"`
}

type ReceiverRefModel struct {
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
}

type RouteModel struct {
	Receiver types.String `tfsdk:"receiver"`
	Match    types.Map    `tfsdk:"match"`
}

type APIAlertPolicy struct {
	Metadata   APIMetadata        `json:"metadata"`
	Spec       APIAlertPolicySpec `json:"spec"`
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

type APIAlertPolicySpec struct {
	Receivers []APIReceiverRef `json:"receivers,omitempty"`
	Routes    []APIRoute       `json:"routes,omitempty"`
}

type APIReceiverRef struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type APIRoute struct {
	Receiver string            `json:"receiver,omitempty"`
	Match    map[string]string `json:"match,omitempty"`
}

func (m *AlertPolicyResourceModel) ToAPIRequest() *APIAlertPolicy {
	spec := APIAlertPolicySpec{}
	if len(m.Receivers) > 0 {
		receivers := make([]APIReceiverRef, len(m.Receivers))
		for i, r := range m.Receivers {
			receivers[i] = APIReceiverRef{Name: r.Name.ValueString(), Namespace: r.Namespace.ValueString()}
		}
		spec.Receivers = receivers
	}
	return &APIAlertPolicy{
		Metadata: APIMetadata{Name: m.Name.ValueString(), Namespace: m.Namespace.ValueString(), Description: m.Description.ValueString()},
		Spec:     spec,
	}
}

func (m *AlertPolicyResourceModel) FromAPIResponse(resp *APIAlertPolicy) {
	m.Name = types.StringValue(resp.Metadata.Name)
	m.Namespace = types.StringValue(resp.Metadata.Namespace)
	m.Description = types.StringValue(resp.Metadata.Description)
	if resp.SystemMeta.UID != "" {
		m.ID = types.StringValue(resp.SystemMeta.UID)
	} else {
		m.ID = types.StringValue(resp.Metadata.Namespace + "/" + resp.Metadata.Name)
	}
}
