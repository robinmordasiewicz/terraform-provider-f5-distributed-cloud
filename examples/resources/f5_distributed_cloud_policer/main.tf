# Example configuration for f5_distributed_cloud_policer

resource "f5_distributed_cloud_policer" "example" {
  name        = "example-policer"
  namespace   = "system"
  description = "Example Policer resource managed by Terraform"

  # Add additional configuration as needed
}
