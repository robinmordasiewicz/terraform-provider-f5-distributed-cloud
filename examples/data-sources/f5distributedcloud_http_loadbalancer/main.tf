# Example configuration for f5distributedcloud_http_loadbalancer data source

data "f5distributedcloud_http_loadbalancer" "example" {
  name      = "existing-http_loadbalancer"
  namespace = "system"
}

output "http_loadbalancer_id" {
  value = data.f5distributedcloud_http_loadbalancer.example.id
}
