// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package validators

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// F5XCNameValidator returns a validator for F5 XC resource names.
// Names must:
// - Start with a lowercase letter
// - Contain only lowercase letters, numbers, and hyphens
// - Not end with a hyphen
// - Be between 1 and 63 characters
func F5XCNameValidator() validator.String {
	return stringvalidator.All(
		stringvalidator.LengthBetween(1, 63),
		stringvalidator.RegexMatches(
			regexp.MustCompile(`^[a-z][a-z0-9-]*[a-z0-9]$|^[a-z]$`),
			"must start with a lowercase letter, contain only lowercase letters, numbers, and hyphens, and not end with a hyphen",
		),
	)
}

// F5XCNamespaceValidator returns a validator specifically for namespace names.
// Namespace names follow the same rules as F5XC names but are limited to 32 characters.
func F5XCNamespaceValidator() validator.String {
	return stringvalidator.All(
		stringvalidator.LengthBetween(1, 32),
		stringvalidator.RegexMatches(
			regexp.MustCompile(`^[a-z][a-z0-9-]*[a-z0-9]$|^[a-z]$`),
			"must start with a lowercase letter, contain only lowercase letters, numbers, and hyphens, and not end with a hyphen (max 32 characters)",
		),
	)
}

// URLValidator returns a validator for URL fields.
func URLValidator() validator.String {
	return stringvalidator.RegexMatches(
		regexp.MustCompile(`^https?://[a-zA-Z0-9][-a-zA-Z0-9]*(\.[a-zA-Z0-9][-a-zA-Z0-9]*)*(:[0-9]+)?(/.*)?$`),
		"must be a valid HTTP or HTTPS URL",
	)
}

// IPAddressValidator returns a validator for IPv4 addresses.
func IPAddressValidator() validator.String {
	return stringvalidator.RegexMatches(
		regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`),
		"must be a valid IPv4 address",
	)
}

// CIDRValidator returns a validator for CIDR notation.
func CIDRValidator() validator.String {
	return stringvalidator.RegexMatches(
		regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)/([0-9]|[1-2][0-9]|3[0-2])$`),
		"must be a valid CIDR notation (e.g., 10.0.0.0/8)",
	)
}

// PortValidator returns a validator for port numbers.
func PortValidator() validator.Int64 {
	return &portValidator{}
}

type portValidator struct{}

func (v *portValidator) Description(ctx context.Context) string {
	return "must be a valid port number between 1 and 65535"
}

func (v *portValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v *portValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueInt64()
	if value < 1 || value > 65535 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Port Number",
			fmt.Sprintf("Port must be between 1 and 65535, got: %d", value),
		)
	}
}

// DNSNameValidator returns a validator for DNS names.
func DNSNameValidator() validator.String {
	return stringvalidator.All(
		stringvalidator.LengthBetween(1, 253),
		stringvalidator.RegexMatches(
			regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`),
			"must be a valid DNS name",
		),
	)
}

// LabelKeyValidator returns a validator for Kubernetes-style label keys.
func LabelKeyValidator() validator.String {
	return stringvalidator.All(
		stringvalidator.LengthBetween(1, 63),
		stringvalidator.RegexMatches(
			regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_.-]*$`),
			"must start with a letter and contain only alphanumeric characters, underscores, dots, and hyphens",
		),
	)
}

// LabelValueValidator returns a validator for Kubernetes-style label values.
func LabelValueValidator() validator.String {
	return stringvalidator.LengthAtMost(63)
}
