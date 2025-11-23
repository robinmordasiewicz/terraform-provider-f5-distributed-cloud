# Example configuration for f5_distributed_cloud_secret_policy data source

data "f5_distributed_cloud_secret_policy" "example" {
  name      = "existing-secret_policy"
  namespace = "system"
}

output "secret_policy_id" {
  value = data.f5_distributed_cloud_secret_policy.example.id
}
