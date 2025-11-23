# Example configuration for f5_distributed_cloud_workload

resource "f5_distributed_cloud_workload" "example" {
  name        = "example-workload"
  namespace   = "system"
  description = "Example Workload resource managed by Terraform"

  # Add additional configuration as needed
}
