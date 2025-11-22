# Example configuration for f5xc_app_api_group data source

data "f5xc_app_api_group" "example" {
  name      = "existing-app_api_group"
  namespace = "system"
}

output "app_api_group_id" {
  value = data.f5xc_app_api_group.example.id
}
