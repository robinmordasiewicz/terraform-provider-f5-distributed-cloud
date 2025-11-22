// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package cloud_site_test

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

// T068: Write acceptance test for cloud_site create (AWS)
func TestAccCloudSiteResource_createAWS(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSiteResourceConfigAWS(rName, "Test AWS Site"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_cloud_site.test", "name", rName),
					resource.TestCheckResourceAttr("f5xc_cloud_site.test", "site_type", "aws_vpc_site"),
					resource.TestCheckResourceAttr("f5xc_cloud_site.test", "cloud_provider", "aws"),
					resource.TestCheckResourceAttrSet("f5xc_cloud_site.test", "id"),
				),
			},
		},
	})
}

// T069: Write acceptance test for cloud_site create (Azure)
func TestAccCloudSiteResource_createAzure(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSiteResourceConfigAzure(rName, "Test Azure Site"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_cloud_site.test", "name", rName),
					resource.TestCheckResourceAttr("f5xc_cloud_site.test", "site_type", "azure_vnet_site"),
					resource.TestCheckResourceAttr("f5xc_cloud_site.test", "cloud_provider", "azure"),
					resource.TestCheckResourceAttrSet("f5xc_cloud_site.test", "id"),
				),
			},
		},
	})
}

// T070: Write acceptance test for cloud_site create (GCP)
func TestAccCloudSiteResource_createGCP(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSiteResourceConfigGCP(rName, "Test GCP Site"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("f5xc_cloud_site.test", "name", rName),
					resource.TestCheckResourceAttr("f5xc_cloud_site.test", "site_type", "gcp_vpc_site"),
					resource.TestCheckResourceAttr("f5xc_cloud_site.test", "cloud_provider", "gcp"),
					resource.TestCheckResourceAttrSet("f5xc_cloud_site.test", "id"),
				),
			},
		},
	})
}

// T071: Write acceptance test for cloud_site update
func TestAccCloudSiteResource_update(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSiteResourceConfigAWS(rName, "Initial"),
				Check:  resource.TestCheckResourceAttr("f5xc_cloud_site.test", "description", "Initial"),
			},
			{
				Config: testAccCloudSiteResourceConfigAWS(rName, "Updated"),
				Check:  resource.TestCheckResourceAttr("f5xc_cloud_site.test", "description", "Updated"),
			},
		},
	})
}

// T072: Write acceptance test for cloud_site delete
func TestAccCloudSiteResource_delete(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSiteResourceConfigAWS(rName, "To be deleted"),
			},
		},
	})
}

// T073: Write acceptance test for cloud_site import
func TestAccCloudSiteResource_import(t *testing.T) {
	rName := "tf-acc-test-" + acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudSiteResourceConfigAWS(rName, "For import"),
			},
			{
				ResourceName:            "f5xc_cloud_site.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"site_type", "cloud_provider"},
				ImportStateId:           rName,
			},
		},
	})
}

func testAccCloudSiteResourceConfigAWS(name, description string) string {
	return fmt.Sprintf(`
provider "f5xc" {}

resource "f5xc_cloud_site" "test" {
  name           = %[1]q
  description    = %[2]q
  site_type      = "aws_vpc_site"
  cloud_provider = "aws"
  region         = "us-east-1"

  vpc_name = "test-vpc"
  cidr     = "10.0.0.0/16"

  subnets {
    name              = "subnet-1"
    cidr              = "10.0.1.0/24"
    availability_zone = "us-east-1a"
    subnet_type       = "public"
  }

  master_nodes {
    instance_type = "t3.xlarge"
    disk_size     = 80
    count         = 1
  }
}
`, name, description)
}

func testAccCloudSiteResourceConfigAzure(name, description string) string {
	return fmt.Sprintf(`
provider "f5xc" {}

resource "f5xc_cloud_site" "test" {
  name           = %[1]q
  description    = %[2]q
  site_type      = "azure_vnet_site"
  cloud_provider = "azure"
  region         = "eastus"

  vpc_name = "test-vnet"
  cidr     = "10.0.0.0/16"

  subnets {
    name              = "subnet-1"
    cidr              = "10.0.1.0/24"
    availability_zone = "1"
    subnet_type       = "public"
  }

  master_nodes {
    instance_type = "Standard_D3_v2"
    disk_size     = 80
    count         = 1
  }
}
`, name, description)
}

func testAccCloudSiteResourceConfigGCP(name, description string) string {
	return fmt.Sprintf(`
provider "f5xc" {}

resource "f5xc_cloud_site" "test" {
  name           = %[1]q
  description    = %[2]q
  site_type      = "gcp_vpc_site"
  cloud_provider = "gcp"
  region         = "us-central1"

  vpc_name = "test-vpc"
  cidr     = "10.0.0.0/16"

  subnets {
    name              = "subnet-1"
    cidr              = "10.0.1.0/24"
    availability_zone = "us-central1-a"
    subnet_type       = "public"
  }

  master_nodes {
    instance_type = "n1-standard-4"
    disk_size     = 80
    count         = 1
  }
}
`, name, description)
}
