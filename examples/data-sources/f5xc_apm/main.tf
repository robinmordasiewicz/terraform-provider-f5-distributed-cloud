# Example configuration for f5xc_apm data source

data "f5xc_apm" "example" {
  name      = "existing-apm"
  namespace = "system"
}

output "apm_id" {
  value = data.f5xc_apm.example.id
}
