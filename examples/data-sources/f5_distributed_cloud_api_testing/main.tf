# Example configuration for f5_distributed_cloud_api_testing data source

data "f5_distributed_cloud_api_testing" "example" {
  name      = "existing-api_testing"
  namespace = "system"
}

output "api_testing_id" {
  value = data.f5_distributed_cloud_api_testing.example.id
}
