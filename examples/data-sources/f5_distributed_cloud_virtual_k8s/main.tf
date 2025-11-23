# Example configuration for f5_distributed_cloud_virtual_k8s data source

data "f5_distributed_cloud_virtual_k8s" "example" {
  name      = "existing-virtual_k8s"
  namespace = "system"
}

output "virtual_k8s_id" {
  value = data.f5_distributed_cloud_virtual_k8s.example.id
}
