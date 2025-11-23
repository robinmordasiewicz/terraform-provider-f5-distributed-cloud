# Example configuration for f5_distributed_cloud_infraprotect_tunnel data source

data "f5_distributed_cloud_infraprotect_tunnel" "example" {
  name      = "existing-infraprotect_tunnel"
  namespace = "system"
}

output "infraprotect_tunnel_id" {
  value = data.f5_distributed_cloud_infraprotect_tunnel.example.id
}
