# Example configuration for f5xc_app_type data source

data "f5xc_app_type" "example" {
  name      = "existing-app_type"
  namespace = "system"
}

output "app_type_id" {
  value = data.f5xc_app_type.example.id
}
