# Example: Healthcheck Data Source
# This example retrieves information about an existing Healthcheck

data "f5xc_healthcheck" "example" {
  name      = "my-healthcheck"
  namespace = "my-namespace"
}

# Output the healthcheck details
output "healthcheck_id" {
  description = "ID of the Healthcheck"
  value       = data.f5xc_healthcheck.example.id
}

output "healthcheck_type" {
  description = "Type of the Healthcheck (HTTP or TCP)"
  value       = data.f5xc_healthcheck.example.healthcheck_type
}

output "healthcheck_interval" {
  description = "Interval between healthcheck requests"
  value       = data.f5xc_healthcheck.example.interval
}

# Use the data source to reference an existing healthcheck
# Healthchecks are commonly used with origin pools
resource "f5xc_origin_pool" "example" {
  name        = "backend-with-healthcheck"
  namespace   = data.f5xc_healthcheck.example.namespace
  description = "Origin pool using ${data.f5xc_healthcheck.example.name} healthcheck"

  origin_servers {
    public_ip {
      ip = "10.0.0.1"
    }
  }

  port               = 8080
  endpoint_selection = "LOCAL_PREFERRED"

  healthcheck {
    name      = data.f5xc_healthcheck.example.name
    namespace = data.f5xc_healthcheck.example.namespace
  }
}
