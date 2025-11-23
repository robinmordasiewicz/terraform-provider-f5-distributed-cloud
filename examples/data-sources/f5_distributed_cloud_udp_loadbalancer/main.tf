# Example configuration for f5_distributed_cloud_udp_loadbalancer data source

data "f5_distributed_cloud_udp_loadbalancer" "example" {
  name      = "existing-udp_loadbalancer"
  namespace = "system"
}

output "udp_loadbalancer_id" {
  value = data.f5_distributed_cloud_udp_loadbalancer.example.id
}
