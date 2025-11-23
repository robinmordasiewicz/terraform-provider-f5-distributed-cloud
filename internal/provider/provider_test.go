// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package provider_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/robinmordasiewicz/terraform-provider-f5distributedcloud/internal/provider"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"f5distributedcloud": providerserver.NewProtocol6WithError(provider.New("test")()),
}

// testAccPreCheck validates the necessary test API keys exist in the testing
// environment.
func testAccPreCheck(t *testing.T) {
	t.Helper()

	apiURL := os.Getenv("F5XC_API_URL")
	if apiURL == "" {
		t.Skip("F5XC_API_URL must be set for acceptance tests")
	}

	// Check for at least one auth method
	apiToken := os.Getenv("F5XC_API_TOKEN")
	p12File := os.Getenv("F5XC_API_P12_FILE")

	if apiToken == "" && p12File == "" {
		t.Skip("Either F5XC_API_TOKEN or F5XC_API_P12_FILE must be set for acceptance tests")
	}
}

// TestAccProvider_tokenAuth tests provider configuration with API token authentication.
// T020: Write acceptance test for token auth
func TestAccProvider_tokenAuth(t *testing.T) {
	// Skip if required env vars not set
	apiURL := os.Getenv("F5XC_API_URL")
	apiToken := os.Getenv("F5XC_API_TOKEN")
	if apiURL == "" || apiToken == "" {
		t.Skip("F5XC_API_URL and F5XC_API_TOKEN must be set for token auth acceptance test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProviderConfig_tokenAuth(apiURL, apiToken),
				Check:  resource.ComposeTestCheckFunc(
				// Provider configured successfully if no error
				),
			},
		},
	})
}

// TestAccProvider_certAuth tests provider configuration with certificate authentication.
// T019: Write acceptance test for certificate auth
func TestAccProvider_certAuth(t *testing.T) {
	// Skip if required env vars not set
	apiURL := os.Getenv("F5XC_API_URL")
	p12File := os.Getenv("F5XC_API_P12_FILE")
	p12Password := os.Getenv("F5XC_API_P12_PASSWORD")
	if apiURL == "" || p12File == "" {
		t.Skip("F5XC_API_URL and F5XC_API_P12_FILE must be set for cert auth acceptance test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProviderConfig_certAuth(apiURL, p12File, p12Password),
				Check:  resource.ComposeTestCheckFunc(
				// Provider configured successfully if no error
				),
			},
		},
	})
}

// TestAccProvider_missingCredentials tests that provider fails with missing credentials.
// T022: Write acceptance test for missing credentials error
func TestAccProvider_missingCredentials(t *testing.T) {
	apiURL := os.Getenv("F5XC_API_URL")
	if apiURL == "" {
		apiURL = "https://test.console.ves.volterra.io/api"
	}

	// Clear auth env vars for this test
	originalToken := os.Getenv("F5XC_API_TOKEN")
	originalP12 := os.Getenv("F5XC_API_P12_FILE")
	os.Unsetenv("F5XC_API_TOKEN")
	os.Unsetenv("F5XC_API_P12_FILE")
	defer func() {
		if originalToken != "" {
			os.Setenv("F5XC_API_TOKEN", originalToken)
		}
		if originalP12 != "" {
			os.Setenv("F5XC_API_P12_FILE", originalP12)
		}
	}()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccProviderConfig_missingCredentials(apiURL),
				ExpectError: regexp.MustCompile(`Missing Authentication Credentials`),
			},
		},
	})
}

// TestAccProvider_invalidCredentials tests that provider handles invalid credentials gracefully.
// T021: Write acceptance test for invalid credentials error handling
func TestAccProvider_invalidCredentials(t *testing.T) {
	apiURL := os.Getenv("F5XC_API_URL")
	if apiURL == "" {
		t.Skip("F5XC_API_URL must be set for invalid credentials test")
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccProviderConfig_tokenAuth(apiURL, "invalid-token-that-will-fail"),
				// Expect 401 error when data source makes API call with invalid credentials
				ExpectError: regexp.MustCompile(`(401|credential invalid|unauthorized)`),
			},
		},
	})
}

// TestAccProvider_missingAPIURL tests that provider fails with missing API URL.
func TestAccProvider_missingAPIURL(t *testing.T) {
	// Clear API URL env var for this test
	originalURL := os.Getenv("F5XC_API_URL")
	os.Unsetenv("F5XC_API_URL")
	defer func() {
		if originalURL != "" {
			os.Setenv("F5XC_API_URL", originalURL)
		}
	}()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccProviderConfig_missingAPIURL(),
				ExpectError: regexp.MustCompile(`Missing API URL`),
			},
		},
	})
}

// testAccProviderConfig_tokenAuth returns a provider configuration for token authentication.
func testAccProviderConfig_tokenAuth(apiURL, apiToken string) string {
	return `
provider "f5distributedcloud" {
  api_url   = "` + apiURL + `"
  api_token = "` + apiToken + `"
}

# Use f5distributedcloud data source to verify provider configuration
data "f5distributedcloud_namespace" "system" {
  name = "system"
}
`
}

// testAccProviderConfig_certAuth returns a provider configuration for certificate authentication.
func testAccProviderConfig_certAuth(apiURL, p12File, p12Password string) string {
	config := `
provider "f5distributedcloud" {
  api_url      = "` + apiURL + `"
  p12_file     = "` + p12File + `"
`
	if p12Password != "" {
		config += `  p12_password = "` + p12Password + `"
`
	}
	// Use a data source to read the "system" namespace (always exists) to verify authentication
	// This tests provider configuration without requiring write permissions
	config += `}

data "f5distributedcloud_namespace" "system" {
  name = "system"
}
`
	return config
}

// testAccProviderConfig_missingCredentials returns a provider configuration without credentials.
func testAccProviderConfig_missingCredentials(apiURL string) string {
	return `
provider "f5distributedcloud" {
  api_url = "` + apiURL + `"
}

# Use f5distributedcloud data source to verify provider configuration
data "f5distributedcloud_namespace" "system" {
  name = "system"
}
`
}

// testAccProviderConfig_missingAPIURL returns a provider configuration without API URL.
func testAccProviderConfig_missingAPIURL() string {
	return `
provider "f5distributedcloud" {
  api_token = "some-token"
}

# Use f5distributedcloud data source to verify provider configuration
data "f5distributedcloud_namespace" "system" {
  name = "system"
}
`
}
