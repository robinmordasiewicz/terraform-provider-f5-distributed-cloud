# Example configuration for f5xc_network_firewall data source

data "f5xc_network_firewall" "example" {
  name      = "existing-network_firewall"
  namespace = "system"
}

output "network_firewall_id" {
  value = data.f5xc_network_firewall.example.id
}
