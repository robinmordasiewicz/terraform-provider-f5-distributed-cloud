# Example configuration for f5_distributed_cloud_proxy data source

data "f5_distributed_cloud_proxy" "example" {
  name      = "existing-proxy"
  namespace = "system"
}

output "proxy_id" {
  value = data.f5_distributed_cloud_proxy.example.id
}
