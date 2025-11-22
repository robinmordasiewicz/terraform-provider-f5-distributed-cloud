# Example configuration for f5xc_bot_infrastructure data source

data "f5xc_bot_infrastructure" "example" {
  name      = "existing-bot_infrastructure"
  namespace = "system"
}

output "bot_infrastructure_id" {
  value = data.f5xc_bot_infrastructure.example.id
}
