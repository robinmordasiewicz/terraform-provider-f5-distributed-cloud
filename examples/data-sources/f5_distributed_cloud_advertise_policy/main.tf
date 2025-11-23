# Example configuration for f5_distributed_cloud_advertise_policy data source

data "f5_distributed_cloud_advertise_policy" "example" {
  name      = "existing-advertise_policy"
  namespace = "system"
}

output "advertise_policy_id" {
  value = data.f5_distributed_cloud_advertise_policy.example.id
}
