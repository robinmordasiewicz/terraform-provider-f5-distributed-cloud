# Example configuration for f5distributedcloud_nat_policy

resource "f5distributedcloud_nat_policy" "example" {
  name        = "example-nat_policy"
  namespace   = "system"
  description = "Example NATPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
