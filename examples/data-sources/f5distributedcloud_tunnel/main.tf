# Example configuration for f5distributedcloud_tunnel data source

data "f5distributedcloud_tunnel" "example" {
  name      = "existing-tunnel"
  namespace = "system"
}

output "tunnel_id" {
  value = data.f5distributedcloud_tunnel.example.id
}
