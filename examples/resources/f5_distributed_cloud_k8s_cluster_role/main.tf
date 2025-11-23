# Example configuration for f5_distributed_cloud_k8s_cluster_role

resource "f5_distributed_cloud_k8s_cluster_role" "example" {
  name        = "example-k8s_cluster_role"
  namespace   = "system"
  description = "Example K8SClusterRole resource managed by Terraform"

  # Add additional configuration as needed
}
