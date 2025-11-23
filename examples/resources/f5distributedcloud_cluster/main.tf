# Example configuration for f5distributedcloud_cluster

resource "f5distributedcloud_cluster" "example" {
  name        = "example-cluster"
  namespace   = "system"
  description = "Example Cluster resource managed by Terraform"

  # Add additional configuration as needed
}
