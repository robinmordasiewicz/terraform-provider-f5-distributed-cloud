// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package namespace_test

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
		t.Skip("Either F5XC_API_TOKEN or F5XC_API_P12_FILE must be set for acceptance tests")
	}
}

// TestAccNamespaceResource_create tests namespace creation.
// T029: Write acceptance test for namespace create
func TestAccNamespaceResource_create(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNamespaceResourceConfig(rName, "Test namespace created by Terraform acceptance test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_namespace.test", "name", rName),
					resource.TestCheckResourceAttr("f5xc_namespace.test", "description", "Test namespace created by Terraform acceptance test"),
					resource.TestCheckResourceAttrSet("f5xc_namespace.test", "id"),
				),
			},
		},
	})
}

// TestAccNamespaceResource_update tests namespace update.
// T030: Write acceptance test for namespace update
func TestAccNamespaceResource_update(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNamespaceResourceConfig(rName, "Initial description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_namespace.test", "description", "Initial description"),
				),
			},
			{
				Config: testAccNamespaceResourceConfig(rName, "Updated description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_namespace.test", "description", "Updated description"),
				),
			},
		},
	})
}

// TestAccNamespaceResource_delete tests namespace deletion.
// T031: Write acceptance test for namespace delete
func TestAccNamespaceResource_delete(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNamespaceResourceConfig(rName, "Namespace to be deleted"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_namespace.test", "name", rName),
				),
			},
			// Destruction is tested automatically after the test
		},
	})
}

// TestAccNamespaceResource_import tests namespace import.
// T032: Write acceptance test for namespace import
func TestAccNamespaceResource_import(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNamespaceResourceConfig(rName, "Test namespace for import"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_namespace.test", "name", rName),
				),
			},
			{
				ResourceName:      "f5xc_namespace.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     rName,
			},
		},
	})
}

func testAccNamespaceResourceConfig(name, description string) string {
	return fmt.Sprintf(`
provider "f5xc" {}

resource "f5xc_namespace" "test" {
  name        = %[1]q
  description = %[2]q
}
`, name, description)
}
