resource "f5xc_network_firewall" "example" {
  name        = "my-network-firewall"
  namespace   = "system"
  description = "Network firewall for perimeter security"
}
