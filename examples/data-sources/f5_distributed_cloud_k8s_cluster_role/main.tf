# Example configuration for f5_distributed_cloud_k8s_cluster_role data source

data "f5_distributed_cloud_k8s_cluster_role" "example" {
  name      = "existing-k8s_cluster_role"
  namespace = "system"
}

output "k8s_cluster_role_id" {
  value = data.f5_distributed_cloud_k8s_cluster_role.example.id
}
