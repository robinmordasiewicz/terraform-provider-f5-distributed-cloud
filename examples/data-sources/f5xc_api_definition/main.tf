# Example configuration for f5xc_api_definition data source

data "f5xc_api_definition" "example" {
  name      = "existing-api_definition"
  namespace = "system"
}

output "api_definition_id" {
  value = data.f5xc_api_definition.example.id
}
