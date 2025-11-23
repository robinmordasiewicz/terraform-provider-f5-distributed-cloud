# Example configuration for f5_distributed_cloud_virtual_site

resource "f5_distributed_cloud_virtual_site" "example" {
  name        = "example-virtual_site"
  namespace   = "system"
  description = "Example VirtualSite resource managed by Terraform"

  # Add additional configuration as needed
}
