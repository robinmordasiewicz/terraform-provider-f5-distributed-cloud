# Example: Look up an existing tunnel
data "f5xc_tunnel" "example" {
  name      = "my-tunnel"
  namespace = "my-namespace"
}

# Output the tunnel type (IPSEC or SSL)
output "tunnel_type" {
  value = data.f5xc_tunnel.example.tunnel_type
}

output "tunnel_id" {
  value = data.f5xc_tunnel.example.id
}
