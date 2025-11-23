# Example configuration for f5distributedcloud_tcp_loadbalancer data source

data "f5distributedcloud_tcp_loadbalancer" "example" {
  name      = "existing-tcp_loadbalancer"
  namespace = "system"
}

output "tcp_loadbalancer_id" {
  value = data.f5distributedcloud_tcp_loadbalancer.example.id
}
