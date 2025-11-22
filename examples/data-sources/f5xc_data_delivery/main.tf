# Example configuration for f5xc_data_delivery data source

data "f5xc_data_delivery" "example" {
  name      = "existing-data_delivery"
  namespace = "system"
}

output "data_delivery_id" {
  value = data.f5xc_data_delivery.example.id
}
