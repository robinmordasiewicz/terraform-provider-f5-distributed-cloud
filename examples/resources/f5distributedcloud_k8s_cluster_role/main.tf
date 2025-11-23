# Example configuration for f5distributedcloud_k8s_cluster_role

resource "f5distributedcloud_k8s_cluster_role" "example" {
  name        = "example-k8s_cluster_role"
  namespace   = "system"
  description = "Example K8SClusterRole resource managed by Terraform"

  # Add additional configuration as needed
}
