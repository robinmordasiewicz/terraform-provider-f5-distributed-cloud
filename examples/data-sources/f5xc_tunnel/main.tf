# Example configuration for f5xc_tunnel data source

data "f5xc_tunnel" "example" {
  name      = "existing-tunnel"
  namespace = "system"
}

output "tunnel_id" {
  value = data.f5xc_tunnel.example.id
}
