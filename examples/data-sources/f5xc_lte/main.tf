# Example configuration for f5xc_lte data source

data "f5xc_lte" "example" {
  name      = "existing-lte"
  namespace = "system"
}

output "lte_id" {
  value = data.f5xc_lte.example.id
}
