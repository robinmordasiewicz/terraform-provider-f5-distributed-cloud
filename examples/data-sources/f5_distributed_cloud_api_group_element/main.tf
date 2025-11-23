# Example configuration for f5_distributed_cloud_api_group_element data source

data "f5_distributed_cloud_api_group_element" "example" {
  name      = "existing-api_group_element"
  namespace = "system"
}

output "api_group_element_id" {
  value = data.f5_distributed_cloud_api_group_element.example.id
}
