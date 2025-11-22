# Example configuration for f5xc_nginx_instance data source

data "f5xc_nginx_instance" "example" {
  name      = "existing-nginx_instance"
  namespace = "system"
}

output "nginx_instance_id" {
  value = data.f5xc_nginx_instance.example.id
}
