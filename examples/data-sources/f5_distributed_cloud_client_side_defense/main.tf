# Example configuration for f5_distributed_cloud_client_side_defense data source

data "f5_distributed_cloud_client_side_defense" "example" {
  name      = "existing-client_side_defense"
  namespace = "system"
}

output "client_side_defense_id" {
  value = data.f5_distributed_cloud_client_side_defense.example.id
}
