# Example configuration for f5xc_healthcheck data source

data "f5xc_healthcheck" "example" {
  name      = "existing-healthcheck"
  namespace = "system"
}

output "healthcheck_id" {
  value = data.f5xc_healthcheck.example.id
}
