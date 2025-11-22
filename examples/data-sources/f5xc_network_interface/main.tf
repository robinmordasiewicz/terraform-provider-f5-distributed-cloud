# Example configuration for f5xc_network_interface data source

data "f5xc_network_interface" "example" {
  name      = "existing-network_interface"
  namespace = "system"
}

output "network_interface_id" {
  value = data.f5xc_network_interface.example.id
}
