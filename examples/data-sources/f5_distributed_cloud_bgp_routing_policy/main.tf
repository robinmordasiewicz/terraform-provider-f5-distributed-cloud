# Example configuration for f5_distributed_cloud_bgp_routing_policy data source

data "f5_distributed_cloud_bgp_routing_policy" "example" {
  name      = "existing-bgp_routing_policy"
  namespace = "system"
}

output "bgp_routing_policy_id" {
  value = data.f5_distributed_cloud_bgp_routing_policy.example.id
}
