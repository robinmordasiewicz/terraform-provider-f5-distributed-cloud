# Example configuration for f5_distributed_cloud_api_definition data source

data "f5_distributed_cloud_api_definition" "example" {
  name      = "existing-api_definition"
  namespace = "system"
}

output "api_definition_id" {
  value = data.f5_distributed_cloud_api_definition.example.id
}
