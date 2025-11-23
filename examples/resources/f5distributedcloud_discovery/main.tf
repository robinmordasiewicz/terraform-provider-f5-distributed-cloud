# Example configuration for f5distributedcloud_discovery

resource "f5distributedcloud_discovery" "example" {
  name        = "example-discovery"
  namespace   = "system"
  description = "Example Discovery resource managed by Terraform"

  # Add additional configuration as needed
}
