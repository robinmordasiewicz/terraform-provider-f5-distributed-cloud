# Example configuration for f5_distributed_cloud_bot_defense_app_infrastructure data source

data "f5_distributed_cloud_bot_defense_app_infrastructure" "example" {
  name      = "existing-bot_defense_app_infrastructure"
  namespace = "system"
}

output "bot_defense_app_infrastructure_id" {
  value = data.f5_distributed_cloud_bot_defense_app_infrastructure.example.id
}
