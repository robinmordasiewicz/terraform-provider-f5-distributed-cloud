# Example configuration for f5xc_virtual_network data source

data "f5xc_virtual_network" "example" {
  name      = "existing-virtual_network"
  namespace = "system"
}

output "virtual_network_id" {
  value = data.f5xc_virtual_network.example.id
}
