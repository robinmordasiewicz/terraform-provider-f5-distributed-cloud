# Example configuration for f5_distributed_cloud_external_connector data source

data "f5_distributed_cloud_external_connector" "example" {
  name      = "existing-external_connector"
  namespace = "system"
}

output "external_connector_id" {
  value = data.f5_distributed_cloud_external_connector.example.id
}
