# Example configuration for f5_distributed_cloud_voltshare_admin_policy data source

data "f5_distributed_cloud_voltshare_admin_policy" "example" {
  name      = "existing-voltshare_admin_policy"
  namespace = "system"
}

output "voltshare_admin_policy_id" {
  value = data.f5_distributed_cloud_voltshare_admin_policy.example.id
}
