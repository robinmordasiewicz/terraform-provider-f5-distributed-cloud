# Example configuration for f5_distributed_cloud_k8s_cluster

resource "f5_distributed_cloud_k8s_cluster" "example" {
  name        = "example-k8s_cluster"
  namespace   = "system"
  description = "Example K8SCluster resource managed by Terraform"

  # Add additional configuration as needed
}
