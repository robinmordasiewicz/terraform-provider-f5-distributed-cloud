# Example configuration for f5xc_address_allocator

resource "f5xc_address_allocator" "example" {
  name        = "example-address_allocator"
  namespace   = "system"
  description = "Example AddressAllocator resource managed by Terraform"

  # Add additional configuration as needed
}
