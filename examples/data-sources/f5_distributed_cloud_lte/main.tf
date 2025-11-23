# Example configuration for f5_distributed_cloud_lte data source

data "f5_distributed_cloud_lte" "example" {
  name      = "existing-lte"
  namespace = "system"
}

output "lte_id" {
  value = data.f5_distributed_cloud_lte.example.id
}
