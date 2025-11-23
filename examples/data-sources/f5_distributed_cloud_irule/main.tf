# Example configuration for f5_distributed_cloud_irule data source

data "f5_distributed_cloud_irule" "example" {
  name      = "existing-irule"
  namespace = "system"
}

output "irule_id" {
  value = data.f5_distributed_cloud_irule.example.id
}
