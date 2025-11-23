# Example configuration for f5_distributed_cloud_virtual_host data source

data "f5_distributed_cloud_virtual_host" "example" {
  name      = "existing-virtual_host"
  namespace = "system"
}

output "virtual_host_id" {
  value = data.f5_distributed_cloud_virtual_host.example.id
}
