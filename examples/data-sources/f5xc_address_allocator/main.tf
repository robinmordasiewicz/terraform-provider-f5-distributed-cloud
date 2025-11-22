# Example configuration for f5xc_address_allocator data source

data "f5xc_address_allocator" "example" {
  name      = "existing-address_allocator"
  namespace = "system"
}

output "address_allocator_id" {
  value = data.f5xc_address_allocator.example.id
}
