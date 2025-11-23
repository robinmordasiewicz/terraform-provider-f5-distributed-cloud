# Example configuration for f5_distributed_cloud_virtual_host

resource "f5_distributed_cloud_virtual_host" "example" {
  name        = "example-virtual_host"
  namespace   = "system"
  description = "Example VirtualHost resource managed by Terraform"

  # Add additional configuration as needed
}
