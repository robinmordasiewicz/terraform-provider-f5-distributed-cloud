# Example configuration for f5_distributed_cloud_network_firewall data source

data "f5_distributed_cloud_network_firewall" "example" {
  name      = "existing-network_firewall"
  namespace = "system"
}

output "network_firewall_id" {
  value = data.f5_distributed_cloud_network_firewall.example.id
}
