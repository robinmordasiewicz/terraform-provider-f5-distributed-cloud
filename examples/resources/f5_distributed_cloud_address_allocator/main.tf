# Example configuration for f5_distributed_cloud_address_allocator

resource "f5_distributed_cloud_address_allocator" "example" {
  name        = "example-address_allocator"
  namespace   = "system"
  description = "Example AddressAllocator resource managed by Terraform"

  # Add additional configuration as needed
}
