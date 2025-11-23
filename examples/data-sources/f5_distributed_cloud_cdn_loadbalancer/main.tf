# Example configuration for f5_distributed_cloud_cdn_loadbalancer data source

data "f5_distributed_cloud_cdn_loadbalancer" "example" {
  name      = "existing-cdn_loadbalancer"
  namespace = "system"
}

output "cdn_loadbalancer_id" {
  value = data.f5_distributed_cloud_cdn_loadbalancer.example.id
}
