# Example configuration for f5_distributed_cloud_k8s_cluster_role_binding data source

data "f5_distributed_cloud_k8s_cluster_role_binding" "example" {
  name      = "existing-k8s_cluster_role_binding"
  namespace = "system"
}

output "k8s_cluster_role_binding_id" {
  value = data.f5_distributed_cloud_k8s_cluster_role_binding.example.id
}
