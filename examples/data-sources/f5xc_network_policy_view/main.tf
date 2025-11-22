# Example configuration for f5xc_network_policy_view data source

data "f5xc_network_policy_view" "example" {
  name      = "existing-network_policy_view"
  namespace = "system"
}

output "network_policy_view_id" {
  value = data.f5xc_network_policy_view.example.id
}
