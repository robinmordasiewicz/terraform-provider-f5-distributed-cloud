# Example configuration for f5xc_udp_loadbalancer data source

data "f5xc_udp_loadbalancer" "example" {
  name      = "existing-udp_loadbalancer"
  namespace = "system"
}

output "udp_loadbalancer_id" {
  value = data.f5xc_udp_loadbalancer.example.id
}
