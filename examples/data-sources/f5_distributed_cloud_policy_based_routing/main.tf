# Example configuration for f5_distributed_cloud_policy_based_routing data source

data "f5_distributed_cloud_policy_based_routing" "example" {
  name      = "existing-policy_based_routing"
  namespace = "system"
}

output "policy_based_routing_id" {
  value = data.f5_distributed_cloud_policy_based_routing.example.id
}
