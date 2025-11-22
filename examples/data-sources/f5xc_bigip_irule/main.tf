# Example configuration for f5xc_bigip_irule data source

data "f5xc_bigip_irule" "example" {
  name      = "existing-bigip_irule"
  namespace = "system"
}

output "bigip_irule_id" {
  value = data.f5xc_bigip_irule.example.id
}
