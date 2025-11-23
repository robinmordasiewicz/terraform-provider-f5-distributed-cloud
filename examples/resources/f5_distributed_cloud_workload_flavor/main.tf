# Example configuration for f5_distributed_cloud_workload_flavor

resource "f5_distributed_cloud_workload_flavor" "example" {
  name        = "example-workload_flavor"
  namespace   = "system"
  description = "Example WorkloadFlavor resource managed by Terraform"

  # Add additional configuration as needed
}
