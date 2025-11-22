# Example configuration for f5xc_infraprotect_asn data source

data "f5xc_infraprotect_asn" "example" {
  name      = "existing-infraprotect_asn"
  namespace = "system"
}

output "infraprotect_asn_id" {
  value = data.f5xc_infraprotect_asn.example.id
}
