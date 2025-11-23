# Example configuration for f5distributedcloud_virtual_site

resource "f5distributedcloud_virtual_site" "example" {
  name        = "example-virtual_site"
  namespace   = "system"
  description = "Example VirtualSite resource managed by Terraform"

  # Add additional configuration as needed
}
