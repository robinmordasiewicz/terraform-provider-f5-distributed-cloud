# Example configuration for f5_distributed_cloud_nat_policy data source

data "f5_distributed_cloud_nat_policy" "example" {
  name      = "existing-nat_policy"
  namespace = "system"
}

output "nat_policy_id" {
  value = data.f5_distributed_cloud_nat_policy.example.id
}
