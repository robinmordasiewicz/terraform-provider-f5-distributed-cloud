# Example configuration for f5_distributed_cloud_bgp_asn_set

resource "f5_distributed_cloud_bgp_asn_set" "example" {
  name        = "example-bgp_asn_set"
  namespace   = "system"
  description = "Example BGPAsnSet resource managed by Terraform"

  # Add additional configuration as needed
}
