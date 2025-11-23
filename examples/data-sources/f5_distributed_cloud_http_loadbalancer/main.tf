# Example configuration for f5_distributed_cloud_http_loadbalancer data source

data "f5_distributed_cloud_http_loadbalancer" "example" {
  name      = "existing-http_loadbalancer"
  namespace = "system"
}

output "http_loadbalancer_id" {
  value = data.f5_distributed_cloud_http_loadbalancer.example.id
}
