# Example configuration for f5distributedcloud_gcp_vpc_site

resource "f5distributedcloud_gcp_vpc_site" "example" {
  name        = "example-gcp_vpc_site"
  namespace   = "system"
  description = "Example GCPVPCSite resource managed by Terraform"

  # Add additional configuration as needed
}
