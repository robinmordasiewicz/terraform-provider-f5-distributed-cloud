# Example configuration for f5_distributed_cloud_bgp_asn_set data source

data "f5_distributed_cloud_bgp_asn_set" "example" {
  name      = "existing-bgp_asn_set"
  namespace = "system"
}

output "bgp_asn_set_id" {
  value = data.f5_distributed_cloud_bgp_asn_set.example.id
}
