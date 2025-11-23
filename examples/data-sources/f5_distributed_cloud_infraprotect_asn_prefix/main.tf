# Example configuration for f5_distributed_cloud_infraprotect_asn_prefix data source

data "f5_distributed_cloud_infraprotect_asn_prefix" "example" {
  name      = "existing-infraprotect_asn_prefix"
  namespace = "system"
}

output "infraprotect_asn_prefix_id" {
  value = data.f5_distributed_cloud_infraprotect_asn_prefix.example.id
}
