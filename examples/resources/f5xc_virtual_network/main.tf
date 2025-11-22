resource "f5xc_virtual_network" "example" {
  name        = "my-vnet"
  namespace   = "system"
  description = "Example virtual network"
}
