# Example configuration for f5_distributed_cloud_nginx_csg data source

data "f5_distributed_cloud_nginx_csg" "example" {
  name      = "existing-nginx_csg"
  namespace = "system"
}

output "nginx_csg_id" {
  value = data.f5_distributed_cloud_nginx_csg.example.id
}
