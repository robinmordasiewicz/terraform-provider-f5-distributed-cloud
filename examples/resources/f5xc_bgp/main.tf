resource "f5xc_bgp" "example" {
  name        = "my-bgp-config"
  namespace   = "system"
  description = "BGP configuration for routing"
  asn         = 65000
}
