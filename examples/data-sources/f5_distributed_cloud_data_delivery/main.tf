# Example configuration for f5_distributed_cloud_data_delivery data source

data "f5_distributed_cloud_data_delivery" "example" {
  name      = "existing-data_delivery"
  namespace = "system"
}

output "data_delivery_id" {
  value = data.f5_distributed_cloud_data_delivery.example.id
}
