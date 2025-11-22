# Example configuration for f5xc_tcp_loadbalancer data source

data "f5xc_tcp_loadbalancer" "example" {
  name      = "existing-tcp_loadbalancer"
  namespace = "system"
}

output "tcp_loadbalancer_id" {
  value = data.f5xc_tcp_loadbalancer.example.id
}
