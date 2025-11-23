# Example configuration for f5_distributed_cloud_nat_policy

resource "f5_distributed_cloud_nat_policy" "example" {
  name        = "example-nat_policy"
  namespace   = "system"
  description = "Example NATPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
