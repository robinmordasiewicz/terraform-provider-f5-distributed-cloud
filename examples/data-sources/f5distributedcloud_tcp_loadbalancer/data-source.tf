# Example: TCP Load Balancer Data Source
# This example retrieves information about an existing TCP Load Balancer

data "f5distributedcloud_tcp_loadbalancer" "example" {
  name      = "my-tcp-lb"
  namespace = "my-namespace"
}

# Output the TCP load balancer details
output "tcp_lb_id" {
  description = "ID of the TCP Load Balancer"
  value       = data.f5distributedcloud_tcp_loadbalancer.example.id
}

output "tcp_lb_description" {
  description = "Description of the TCP Load Balancer"
  value       = data.f5distributedcloud_tcp_loadbalancer.example.description
}

output "tcp_lb_port" {
  description = "Listen port of the TCP Load Balancer"
  value       = data.f5distributedcloud_tcp_loadbalancer.example.listen_port
}

# Use the data source to reference an existing TCP load balancer
# This is useful when you want to create related resources
resource "f5distributedcloud_origin_pool" "backend" {
  name        = "tcp-backend-pool"
  namespace   = data.f5distributedcloud_tcp_loadbalancer.example.namespace
  description = "Backend pool for ${data.f5distributedcloud_tcp_loadbalancer.example.name}"

  origin_servers {
    public_ip {
      ip = "10.0.0.1"
    }
  }

  port               = 3306
  endpoint_selection = "LOCAL_PREFERRED"
}
