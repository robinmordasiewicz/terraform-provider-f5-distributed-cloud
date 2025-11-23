# Example configuration for f5distributedcloud_healthcheck data source

data "f5distributedcloud_healthcheck" "example" {
  name      = "existing-healthcheck"
  namespace = "system"
}

output "healthcheck_id" {
  value = data.f5distributedcloud_healthcheck.example.id
}
