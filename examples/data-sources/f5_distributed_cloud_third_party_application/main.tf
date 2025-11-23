# Example configuration for f5_distributed_cloud_third_party_application data source

data "f5_distributed_cloud_third_party_application" "example" {
  name      = "existing-third_party_application"
  namespace = "system"
}

output "third_party_application_id" {
  value = data.f5_distributed_cloud_third_party_application.example.id
}
