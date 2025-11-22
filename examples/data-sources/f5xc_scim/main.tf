# Example configuration for f5xc_scim data source

data "f5xc_scim" "example" {
  name      = "existing-scim"
  namespace = "system"
}

output "scim_id" {
  value = data.f5xc_scim.example.id
}
