# Example configuration for f5_distributed_cloud_rbac_policy data source

data "f5_distributed_cloud_rbac_policy" "example" {
  name      = "existing-rbac_policy"
  namespace = "system"
}

output "rbac_policy_id" {
  value = data.f5_distributed_cloud_rbac_policy.example.id
}
