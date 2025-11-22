# Example configuration for f5xc_endpoint data source

data "f5xc_endpoint" "example" {
  name      = "existing-endpoint"
  namespace = "system"
}

output "endpoint_id" {
  value = data.f5xc_endpoint.example.id
}
