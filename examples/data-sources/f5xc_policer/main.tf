# Example configuration for f5xc_policer data source

data "f5xc_policer" "example" {
  name      = "existing-policer"
  namespace = "system"
}

output "policer_id" {
  value = data.f5xc_policer.example.id
}
