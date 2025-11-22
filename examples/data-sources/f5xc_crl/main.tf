# Example configuration for f5xc_crl data source

data "f5xc_crl" "example" {
  name      = "existing-crl"
  namespace = "system"
}

output "crl_id" {
  value = data.f5xc_crl.example.id
}
