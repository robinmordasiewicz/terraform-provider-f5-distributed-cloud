# Example configuration for f5xc_proxy data source

data "f5xc_proxy" "example" {
  name      = "existing-proxy"
  namespace = "system"
}

output "proxy_id" {
  value = data.f5xc_proxy.example.id
}
