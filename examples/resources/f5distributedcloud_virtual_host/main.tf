# Example configuration for f5distributedcloud_virtual_host

resource "f5distributedcloud_virtual_host" "example" {
  name        = "example-virtual_host"
  namespace   = "system"
  description = "Example VirtualHost resource managed by Terraform"

  # Add additional configuration as needed
}
