# Example configuration for f5_distributed_cloud_k8s_cluster_role_binding

resource "f5_distributed_cloud_k8s_cluster_role_binding" "example" {
  name        = "example-k8s_cluster_role_binding"
  namespace   = "system"
  description = "Example K8SClusterRoleBinding resource managed by Terraform"

  # Add additional configuration as needed
}
