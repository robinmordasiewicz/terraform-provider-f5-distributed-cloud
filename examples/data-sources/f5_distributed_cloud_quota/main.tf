# Example configuration for f5_distributed_cloud_quota data source

data "f5_distributed_cloud_quota" "example" {
  name      = "existing-quota"
  namespace = "system"
}

output "quota_id" {
  value = data.f5_distributed_cloud_quota.example.id
}
