# Example configuration for f5_distributed_cloud_voltstack_site

resource "f5_distributed_cloud_voltstack_site" "example" {
  name        = "example-voltstack_site"
  namespace   = "system"
  description = "Example VoltstackSite resource managed by Terraform"

  # Add additional configuration as needed
}
