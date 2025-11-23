# Example configuration for f5distributedcloud_workload

resource "f5distributedcloud_workload" "example" {
  name        = "example-workload"
  namespace   = "system"
  description = "Example Workload resource managed by Terraform"

  # Add additional configuration as needed
}
