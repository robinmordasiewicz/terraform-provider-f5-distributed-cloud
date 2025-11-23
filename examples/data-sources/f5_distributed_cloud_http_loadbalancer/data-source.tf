# Example: HTTP Load Balancer Data Source
# This example retrieves information about an existing HTTP Load Balancer

data "f5_distributed_cloud_http_loadbalancer" "example" {
  name      = "my-load-balancer"
  namespace = "my-namespace"
}

# Output the load balancer details
output "lb_id" {
  description = "ID of the HTTP Load Balancer"
  value       = data.f5_distributed_cloud_http_loadbalancer.example.id
}

output "lb_description" {
  description = "Description of the HTTP Load Balancer"
  value       = data.f5_distributed_cloud_http_loadbalancer.example.description
}

output "lb_domains" {
  description = "Domains configured on the HTTP Load Balancer"
  value       = data.f5_distributed_cloud_http_loadbalancer.example.domains
}

# Use the data source to reference an existing load balancer
# This is useful for importing existing resources or creating dependencies
resource "f5_distributed_cloud_origin_pool" "backend" {
  name        = "backend-pool"
  namespace   = data.f5_distributed_cloud_http_loadbalancer.example.namespace
  description = "Backend pool for ${data.f5_distributed_cloud_http_loadbalancer.example.name}"

  origin_servers {
    public_ip {
      ip = "10.0.0.1"
    }
  }

  port             = 8080
  endpoint_selection = "LOCAL_PREFERRED"
}
