# Example configuration for f5_distributed_cloud_address_allocator data source

data "f5_distributed_cloud_address_allocator" "example" {
  name      = "existing-address_allocator"
  namespace = "system"
}

output "address_allocator_id" {
  value = data.f5_distributed_cloud_address_allocator.example.id
}
