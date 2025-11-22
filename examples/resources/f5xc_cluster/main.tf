# Example configuration for f5xc_cluster

resource "f5xc_cluster" "example" {
  name        = "example-cluster"
  namespace   = "system"
  description = "Example Cluster resource managed by Terraform"

  # Add additional configuration as needed
}
