# Example configuration for f5_distributed_cloud_nginx_server data source

data "f5_distributed_cloud_nginx_server" "example" {
  name      = "existing-nginx_server"
  namespace = "system"
}

output "nginx_server_id" {
  value = data.f5_distributed_cloud_nginx_server.example.id
}
