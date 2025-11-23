# Example configuration for f5_distributed_cloud_user_identification data source

data "f5_distributed_cloud_user_identification" "example" {
  name      = "existing-user_identification"
  namespace = "system"
}

output "user_identification_id" {
  value = data.f5_distributed_cloud_user_identification.example.id
}
