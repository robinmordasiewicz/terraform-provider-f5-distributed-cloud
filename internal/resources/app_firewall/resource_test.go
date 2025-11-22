// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package app_firewall_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/provider"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"f5xc": providerserver.NewProtocol6WithError(provider.New("test")()),
}

func testAccPreCheck(t *testing.T) {
	t.Helper()
	if os.Getenv("F5XC_API_URL") == "" {
		t.Skip("F5XC_API_URL must be set for acceptance tests")
	}
	if os.Getenv("F5XC_API_TOKEN") == "" && os.Getenv("F5XC_API_P12_FILE") == "" {
		t.Skip("Either F5XC_API_TOKEN or F5XC_API_P12_FILE must be set")
	}
}

// T087: Write acceptance test for app_firewall create
func TestAccAppFirewallResource_create(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)
	namespace := os.Getenv("F5XC_TEST_NAMESPACE")
	if namespace == "" {
		namespace = "system"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallResourceConfig(rName, namespace, "Test App Firewall", "monitoring"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_app_firewall.test", "name", rName),
					resource.TestCheckResourceAttr("f5xc_app_firewall.test", "namespace", namespace),
					resource.TestCheckResourceAttr("f5xc_app_firewall.test", "mode", "monitoring"),
					resource.TestCheckResourceAttrSet("f5xc_app_firewall.test", "id"),
				),
			},
		},
	})
}

// T088: Write acceptance test for app_firewall update
func TestAccAppFirewallResource_update(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)
	namespace := os.Getenv("F5XC_TEST_NAMESPACE")
	if namespace == "" {
		namespace = "system"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallResourceConfig(rName, namespace, "Initial", "monitoring"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_app_firewall.test", "description", "Initial"),
					resource.TestCheckResourceAttr("f5xc_app_firewall.test", "mode", "monitoring"),
				),
			},
			{
				Config: testAccAppFirewallResourceConfig(rName, namespace, "Updated", "blocking"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_app_firewall.test", "description", "Updated"),
					resource.TestCheckResourceAttr("f5xc_app_firewall.test", "mode", "blocking"),
				),
			},
		},
	})
}

// T089: Write acceptance test for app_firewall delete
func TestAccAppFirewallResource_delete(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)
	namespace := os.Getenv("F5XC_TEST_NAMESPACE")
	if namespace == "" {
		namespace = "system"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallResourceConfig(rName, namespace, "To be deleted", "monitoring"),
			},
		},
	})
}

// T090: Write acceptance test for app_firewall import
func TestAccAppFirewallResource_import(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)
	namespace := os.Getenv("F5XC_TEST_NAMESPACE")
	if namespace == "" {
		namespace = "system"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallResourceConfig(rName, namespace, "For import", "monitoring"),
			},
			{
				ResourceName:      "f5xc_app_firewall.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     namespace + "/" + rName,
			},
		},
	})
}

// T091: Write acceptance test for app_firewall with bot protection
func TestAccAppFirewallResource_botProtection(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)
	namespace := os.Getenv("F5XC_TEST_NAMESPACE")
	if namespace == "" {
		namespace = "system"
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppFirewallResourceConfigWithBotProtection(rName, namespace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_app_firewall.test", "name", rName),
					resource.TestCheckResourceAttr("f5xc_app_firewall.test", "bot_protection.good_bot_action", "allow"),
					resource.TestCheckResourceAttr("f5xc_app_firewall.test", "bot_protection.malicious_bot_action", "block"),
				),
			},
		},
	})
}

func testAccAppFirewallResourceConfig(name, namespace, description, mode string) string {
	return fmt.Sprintf(`
provider "f5xc" {}

resource "f5xc_app_firewall" "test" {
  name        = %[1]q
  namespace   = %[2]q
  description = %[3]q
  mode        = %[4]q

  detection_settings {
    signature_based_bot_protection = true
    violation_rating               = "medium"
  }
}
`, name, namespace, description, mode)
}

func testAccAppFirewallResourceConfigWithBotProtection(name, namespace string) string {
	return fmt.Sprintf(`
provider "f5xc" {}

resource "f5xc_app_firewall" "test" {
  name        = %[1]q
  namespace   = %[2]q
  description = "App Firewall with Bot Protection"
  mode        = "blocking"

  detection_settings {
    signature_based_bot_protection = true
    xhr_check                      = true
    cookie_protection              = true
    violation_rating               = "high"
  }

  bot_protection {
    good_bot_action       = "allow"
    malicious_bot_action  = "block"
    suspicious_bot_action = "challenge"
  }
}
`, name, namespace)
}
