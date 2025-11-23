# Example configuration for f5_distributed_cloud_fleet

resource "f5_distributed_cloud_fleet" "example" {
  name        = "example-fleet"
  namespace   = "system"
  description = "Example Fleet resource managed by Terraform"

  # Add additional configuration as needed
}
