# Example configuration for f5distributedcloud_k8s_cluster

resource "f5distributedcloud_k8s_cluster" "example" {
  name        = "example-k8s_cluster"
  namespace   = "system"
  description = "Example K8SCluster resource managed by Terraform"

  # Add additional configuration as needed
}
