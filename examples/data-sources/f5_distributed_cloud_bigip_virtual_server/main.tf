# Example configuration for f5_distributed_cloud_bigip_virtual_server data source

data "f5_distributed_cloud_bigip_virtual_server" "example" {
  name      = "existing-bigip_virtual_server"
  namespace = "system"
}

output "bigip_virtual_server_id" {
  value = data.f5_distributed_cloud_bigip_virtual_server.example.id
}
