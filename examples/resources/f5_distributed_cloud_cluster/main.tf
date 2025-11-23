# Example configuration for f5_distributed_cloud_cluster

resource "f5_distributed_cloud_cluster" "example" {
  name        = "example-cluster"
  namespace   = "system"
  description = "Example Cluster resource managed by Terraform"

  # Add additional configuration as needed
}
