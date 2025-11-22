# Example: Look up an existing API definition
data "f5xc_api_definition" "example" {
  name      = "my-api-definition"
  namespace = "my-namespace"
}

# Output the swagger URL
output "swagger_url" {
  value = data.f5xc_api_definition.example.swagger_url
}

output "api_definition_id" {
  value = data.f5xc_api_definition.example.id
}
