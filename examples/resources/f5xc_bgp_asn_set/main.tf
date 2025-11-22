resource "f5xc_bgp_asn_set" "example" {
  name        = "my-asn-set"
  namespace   = "system"
  description = "BGP ASN set for routing policies"
  asns        = [65000, 65001, 65002]
}
