// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// NameValidator returns a validator for F5 XC resource names.
// Names must be lowercase, start with a letter, and contain only letters, numbers, and hyphens.
func NameValidator() validator.String {
	return stringvalidator.All(
		stringvalidator.LengthBetween(1, 64),
		stringvalidator.RegexMatches(
			regexp.MustCompile(`^[a-z][a-z0-9-]*[a-z0-9]$|^[a-z]$`),
			"must be lowercase, start with a letter, end with a letter or number, and contain only letters, numbers, and hyphens",
		),
	)
}

// NamespaceValidator returns a validator for namespace names.
func NamespaceValidator() validator.String {
	return NameValidator()
}

// DomainValidator returns a validator for domain names.
func DomainValidator() validator.String {
	return &domainValidator{}
}

type domainValidator struct{}

func (v domainValidator) Description(ctx context.Context) string {
	return "must be a valid domain name"
}

func (v domainValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v domainValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	// Basic domain name validation
	if len(value) > 253 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Domain Name",
			"Domain name must be 253 characters or less",
		)
		return
	}

	// Check for valid characters and structure
	domainRegex := regexp.MustCompile(`^(\*\.)?([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)*[a-zA-Z]{2,}$`)
	if !domainRegex.MatchString(value) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Domain Name",
			fmt.Sprintf("%q is not a valid domain name", value),
		)
	}
}

// URLValidator returns a validator for URLs.
func URLValidator() validator.String {
	return &urlValidator{}
}

type urlValidator struct{}

func (v urlValidator) Description(ctx context.Context) string {
	return "must be a valid URL"
}

func (v urlValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v urlValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	parsedURL, err := url.Parse(value)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid URL",
			fmt.Sprintf("failed to parse URL: %s", err),
		)
		return
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid URL",
			"URL scheme must be http or https",
		)
	}

	if parsedURL.Host == "" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid URL",
			"URL must include a host",
		)
	}
}

// IPAddressValidator returns a validator for IP addresses.
func IPAddressValidator() validator.String {
	return &ipAddressValidator{}
}

type ipAddressValidator struct{}

func (v ipAddressValidator) Description(ctx context.Context) string {
	return "must be a valid IP address"
}

func (v ipAddressValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ipAddressValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if net.ParseIP(value) == nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid IP Address",
			fmt.Sprintf("%q is not a valid IP address", value),
		)
	}
}

// CIDRValidator returns a validator for CIDR blocks.
func CIDRValidator() validator.String {
	return &cidrValidator{}
}

type cidrValidator struct{}

func (v cidrValidator) Description(ctx context.Context) string {
	return "must be a valid CIDR block"
}

func (v cidrValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v cidrValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	_, _, err := net.ParseCIDR(value)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid CIDR Block",
			fmt.Sprintf("%q is not a valid CIDR block: %s", value, err),
		)
	}
}

// PortValidator returns a validator for port numbers.
func PortValidator() validator.String {
	return stringvalidator.RegexMatches(
		regexp.MustCompile(`^([1-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$`),
		"must be a valid port number (1-65535)",
	)
}

// OneOfValidator returns a validator that ensures a value is one of the specified options.
func OneOfValidator(options ...string) validator.String {
	return stringvalidator.OneOf(options...)
}

// NotEmptyValidator returns a validator that ensures a string is not empty or whitespace-only.
func NotEmptyValidator() validator.String {
	return &notEmptyValidator{}
}

type notEmptyValidator struct{}

func (v notEmptyValidator) Description(ctx context.Context) string {
	return "must not be empty or whitespace-only"
}

func (v notEmptyValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v notEmptyValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if strings.TrimSpace(value) == "" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Value",
			"value must not be empty or whitespace-only",
		)
	}
}
