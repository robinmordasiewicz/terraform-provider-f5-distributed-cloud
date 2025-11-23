# Example configuration for f5_distributed_cloud_tunnel data source

data "f5_distributed_cloud_tunnel" "example" {
  name      = "existing-tunnel"
  namespace = "system"
}

output "tunnel_id" {
  value = data.f5_distributed_cloud_tunnel.example.id
}
