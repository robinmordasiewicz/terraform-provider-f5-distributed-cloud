# Example configuration for f5xc_infraprotect_asn_prefix data source

data "f5xc_infraprotect_asn_prefix" "example" {
  name      = "existing-infraprotect_asn_prefix"
  namespace = "system"
}

output "infraprotect_asn_prefix_id" {
  value = data.f5xc_infraprotect_asn_prefix.example.id
}
