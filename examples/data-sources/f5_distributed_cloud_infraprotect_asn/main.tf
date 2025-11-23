# Example configuration for f5_distributed_cloud_infraprotect_asn data source

data "f5_distributed_cloud_infraprotect_asn" "example" {
  name      = "existing-infraprotect_asn"
  namespace = "system"
}

output "infraprotect_asn_id" {
  value = data.f5_distributed_cloud_infraprotect_asn.example.id
}
