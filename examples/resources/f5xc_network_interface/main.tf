resource "f5xc_network_interface" "example" {
  name        = "my-network-interface"
  namespace   = "system"
  description = "Site network interface configuration"
}
