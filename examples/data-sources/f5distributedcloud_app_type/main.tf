# Example configuration for f5distributedcloud_app_type data source

data "f5distributedcloud_app_type" "example" {
  name      = "existing-app_type"
  namespace = "system"
}

output "app_type_id" {
  value = data.f5distributedcloud_app_type.example.id
}
