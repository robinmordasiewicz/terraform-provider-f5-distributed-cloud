# Example configuration for f5xc_data_group data source

data "f5xc_data_group" "example" {
  name      = "existing-data_group"
  namespace = "system"
}

output "data_group_id" {
  value = data.f5xc_data_group.example.id
}
