# Example: Application Firewall Data Source
# This example retrieves information about an existing Application Firewall

data "f5distributedcloud_app_firewall" "example" {
  name      = "my-waf-policy"
  namespace = "my-namespace"
}

# Output the application firewall details
output "waf_id" {
  description = "ID of the Application Firewall"
  value       = data.f5distributedcloud_app_firewall.example.id
}

output "waf_description" {
  description = "Description of the Application Firewall"
  value       = data.f5distributedcloud_app_firewall.example.description
}

output "waf_mode" {
  description = "Mode of the Application Firewall (BLOCK or DETECT)"
  value       = data.f5distributedcloud_app_firewall.example.mode
}

# Use the data source to reference an existing WAF policy
# This is useful when you want to attach an existing WAF to a new load balancer
resource "f5distributedcloud_http_loadbalancer" "protected" {
  name        = "protected-app"
  namespace   = data.f5distributedcloud_app_firewall.example.namespace
  description = "Load balancer protected by ${data.f5distributedcloud_app_firewall.example.name}"

  domains = ["app.example.com"]

  http {
    port = 80
  }

  default_route_pools {
    pool {
      name      = "backend-pool"
      namespace = data.f5distributedcloud_app_firewall.example.namespace
    }
    weight = 1
  }

  app_firewall {
    name      = data.f5distributedcloud_app_firewall.example.name
    namespace = data.f5distributedcloud_app_firewall.example.namespace
  }
}
