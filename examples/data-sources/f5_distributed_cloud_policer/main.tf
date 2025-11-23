# Example configuration for f5_distributed_cloud_policer data source

data "f5_distributed_cloud_policer" "example" {
  name      = "existing-policer"
  namespace = "system"
}

output "policer_id" {
  value = data.f5_distributed_cloud_policer.example.id
}
