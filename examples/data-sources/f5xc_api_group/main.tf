# Example configuration for f5xc_api_group data source

data "f5xc_api_group" "example" {
  name      = "existing-api_group"
  namespace = "system"
}

output "api_group_id" {
  value = data.f5xc_api_group.example.id
}
