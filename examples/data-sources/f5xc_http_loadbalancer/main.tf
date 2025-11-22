# Example configuration for f5xc_http_loadbalancer data source

data "f5xc_http_loadbalancer" "example" {
  name      = "existing-http_loadbalancer"
  namespace = "system"
}

output "http_loadbalancer_id" {
  value = data.f5xc_http_loadbalancer.example.id
}
