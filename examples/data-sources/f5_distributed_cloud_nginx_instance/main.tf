# Example configuration for f5_distributed_cloud_nginx_instance data source

data "f5_distributed_cloud_nginx_instance" "example" {
  name      = "existing-nginx_instance"
  namespace = "system"
}

output "nginx_instance_id" {
  value = data.f5_distributed_cloud_nginx_instance.example.id
}
