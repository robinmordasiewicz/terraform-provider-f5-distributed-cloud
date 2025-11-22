# Example configuration for f5xc_irule data source

data "f5xc_irule" "example" {
  name      = "existing-irule"
  namespace = "system"
}

output "irule_id" {
  value = data.f5xc_irule.example.id
}
