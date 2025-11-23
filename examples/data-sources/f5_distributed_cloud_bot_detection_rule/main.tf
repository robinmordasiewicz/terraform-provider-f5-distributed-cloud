# Example configuration for f5_distributed_cloud_bot_detection_rule data source

data "f5_distributed_cloud_bot_detection_rule" "example" {
  name      = "existing-bot_detection_rule"
  namespace = "system"
}

output "bot_detection_rule_id" {
  value = data.f5_distributed_cloud_bot_detection_rule.example.id
}
