# Example: Service Policy Data Source
# This example retrieves information about an existing Service Policy

data "f5xc_service_policy" "example" {
  name      = "my-service-policy"
  namespace = "my-namespace"
}

# Output the service policy details
output "policy_id" {
  description = "ID of the Service Policy"
  value       = data.f5xc_service_policy.example.id
}

output "policy_description" {
  description = "Description of the Service Policy"
  value       = data.f5xc_service_policy.example.description
}

output "policy_algo" {
  description = "Algorithm used by the Service Policy"
  value       = data.f5xc_service_policy.example.algo
}

# Use the data source to reference an existing service policy
# Service policies are commonly used to control traffic flow and access
resource "f5xc_http_loadbalancer" "protected" {
  name        = "protected-app"
  namespace   = data.f5xc_service_policy.example.namespace
  description = "Load balancer protected by ${data.f5xc_service_policy.example.name}"

  domains = ["secure.example.com"]

  http {
    port = 80
  }

  default_route_pools {
    pool {
      name      = "backend-pool"
      namespace = data.f5xc_service_policy.example.namespace
    }
    weight = 1
  }

  service_policy {
    name      = data.f5xc_service_policy.example.name
    namespace = data.f5xc_service_policy.example.namespace
  }
}
