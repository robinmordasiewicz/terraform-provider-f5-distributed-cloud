// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package http_loadbalancer_test

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

// TestAccHTTPLoadBalancerResource_create tests HTTP Load Balancer creation.
// T042: Write acceptance test for http_loadbalancer create
func TestAccHTTPLoadBalancerResource_create(t *testing.T) {
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
				Config: testAccHTTPLoadBalancerResourceConfig(rName, namespace, "Test HTTP LB"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5_distributed_cloud_http_loadbalancer.test", "name", rName),
					resource.TestCheckResourceAttr("f5_distributed_cloud_http_loadbalancer.test", "namespace", namespace),
					resource.TestCheckResourceAttrSet("f5_distributed_cloud_http_loadbalancer.test", "id"),
				),
			},
		},
	})
}

// TestAccHTTPLoadBalancerResource_update tests HTTP Load Balancer update.
// T043: Write acceptance test for http_loadbalancer update
func TestAccHTTPLoadBalancerResource_update(t *testing.T) {
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
				Config: testAccHTTPLoadBalancerResourceConfig(rName, namespace, "Initial description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5_distributed_cloud_http_loadbalancer.test", "description", "Initial description"),
				),
			},
			{
				Config: testAccHTTPLoadBalancerResourceConfig(rName, namespace, "Updated description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5_distributed_cloud_http_loadbalancer.test", "description", "Updated description"),
				),
			},
		},
	})
}

// TestAccHTTPLoadBalancerResource_delete tests HTTP Load Balancer deletion.
// T044: Write acceptance test for http_loadbalancer delete
func TestAccHTTPLoadBalancerResource_delete(t *testing.T) {
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
				Config: testAccHTTPLoadBalancerResourceConfig(rName, namespace, "To be deleted"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5_distributed_cloud_http_loadbalancer.test", "name", rName),
				),
			},
		},
	})
}

// TestAccHTTPLoadBalancerResource_import tests HTTP Load Balancer import.
// T045: Write acceptance test for http_loadbalancer import
func TestAccHTTPLoadBalancerResource_import(t *testing.T) {
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
				Config: testAccHTTPLoadBalancerResourceConfig(rName, namespace, "For import test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5_distributed_cloud_http_loadbalancer.test", "name", rName),
				),
			},
			{
				ResourceName:      "f5_distributed_cloud_http_loadbalancer.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     namespace + "/" + rName,
			},
		},
	})
}

func testAccHTTPLoadBalancerResourceConfig(name, namespace, description string) string {
	return fmt.Sprintf(`
provider "f5_distributed_cloud" {}

resource "f5_distributed_cloud_http_loadbalancer" "test" {
  name        = %[1]q
  namespace   = %[2]q
  description = %[3]q

  domains = ["test.example.com"]

  http_type      = "http"
  advertise_port = 80
  disable_waf    = true
}
`, name, namespace, description)
}
