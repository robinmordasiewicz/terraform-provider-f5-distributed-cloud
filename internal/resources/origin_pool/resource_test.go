// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package origin_pool_test

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
	"f5_distributed_cloud": providerserver.NewProtocol6WithError(provider.New("test")()),
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

// T055: Write acceptance test for origin_pool create
func TestAccOriginPoolResource_create(t *testing.T) {
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
				Config: testAccOriginPoolResourceConfig(rName, namespace, "Test Origin Pool"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5_distributed_cloud_origin_pool.test", "name", rName),
					resource.TestCheckResourceAttr("f5_distributed_cloud_origin_pool.test", "namespace", namespace),
					resource.TestCheckResourceAttrSet("f5_distributed_cloud_origin_pool.test", "id"),
				),
			},
		},
	})
}

// T056: Write acceptance test for origin_pool update
func TestAccOriginPoolResource_update(t *testing.T) {
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
				Config: testAccOriginPoolResourceConfig(rName, namespace, "Initial"),
				Check:  resource.TestCheckResourceAttr("f5_distributed_cloud_origin_pool.test", "description", "Initial"),
			},
			{
				Config: testAccOriginPoolResourceConfig(rName, namespace, "Updated"),
				Check:  resource.TestCheckResourceAttr("f5_distributed_cloud_origin_pool.test", "description", "Updated"),
			},
		},
	})
}

// T057: Write acceptance test for origin_pool delete
func TestAccOriginPoolResource_delete(t *testing.T) {
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
				Config: testAccOriginPoolResourceConfig(rName, namespace, "To be deleted"),
			},
		},
	})
}

// T058: Write acceptance test for origin_pool import
func TestAccOriginPoolResource_import(t *testing.T) {
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
				Config: testAccOriginPoolResourceConfig(rName, namespace, "For import"),
			},
			{
				ResourceName:      "f5_distributed_cloud_origin_pool.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     namespace + "/" + rName,
			},
		},
	})
}

func testAccOriginPoolResourceConfig(name, namespace, description string) string {
	return fmt.Sprintf(`
provider "f5_distributed_cloud" {}

resource "f5_distributed_cloud_origin_pool" "test" {
  name        = %[1]q
  namespace   = %[2]q
  description = %[3]q

  port              = 8080
  endpoint_protocol = "http"

  origin_servers {
    public_name {
      dns_name = "backend.example.com"
    }
  }
}
`, name, namespace, description)
}
