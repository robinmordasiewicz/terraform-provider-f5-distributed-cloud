# Example configuration for f5_distributed_cloud_namespace data source

data "f5_distributed_cloud_namespace" "example" {
  name      = "existing-namespace"
  namespace = "system"
}

output "namespace_id" {
  value = data.f5_distributed_cloud_namespace.example.id
}
