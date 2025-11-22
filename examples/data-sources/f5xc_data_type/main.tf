# Example configuration for f5xc_data_type data source

data "f5xc_data_type" "example" {
  name      = "existing-data_type"
  namespace = "system"
}

output "data_type_id" {
  value = data.f5xc_data_type.example.id
}
