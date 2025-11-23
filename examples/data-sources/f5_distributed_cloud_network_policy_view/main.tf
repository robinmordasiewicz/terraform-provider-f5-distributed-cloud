# Example configuration for f5_distributed_cloud_network_policy_view data source

data "f5_distributed_cloud_network_policy_view" "example" {
  name      = "existing-network_policy_view"
  namespace = "system"
}

output "network_policy_view_id" {
  value = data.f5_distributed_cloud_network_policy_view.example.id
}
