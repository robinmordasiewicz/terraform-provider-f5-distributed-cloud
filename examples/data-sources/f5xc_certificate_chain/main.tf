# Example configuration for f5xc_certificate_chain data source

data "f5xc_certificate_chain" "example" {
  name      = "existing-certificate_chain"
  namespace = "system"
}

output "certificate_chain_id" {
  value = data.f5xc_certificate_chain.example.id
}
