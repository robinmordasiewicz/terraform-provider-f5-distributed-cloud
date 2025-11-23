# Example configuration for f5_distributed_cloud_virtual_network data source

data "f5_distributed_cloud_virtual_network" "example" {
  name      = "existing-virtual_network"
  namespace = "system"
}

output "virtual_network_id" {
  value = data.f5_distributed_cloud_virtual_network.example.id
}
