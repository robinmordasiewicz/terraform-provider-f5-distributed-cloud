# Example configuration for f5_distributed_cloud_healthcheck data source

data "f5_distributed_cloud_healthcheck" "example" {
  name      = "existing-healthcheck"
  namespace = "system"
}

output "healthcheck_id" {
  value = data.f5_distributed_cloud_healthcheck.example.id
}
