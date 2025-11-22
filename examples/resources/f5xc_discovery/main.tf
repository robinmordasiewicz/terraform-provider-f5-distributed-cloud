# Example configuration for f5xc_discovery

resource "f5xc_discovery" "example" {
  name        = "example-discovery"
  namespace   = "system"
  description = "Example Discovery resource managed by Terraform"

  # Add additional configuration as needed
}
