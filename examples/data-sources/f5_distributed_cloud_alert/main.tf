# Example configuration for f5_distributed_cloud_alert data source

data "f5_distributed_cloud_alert" "example" {
  name      = "existing-alert"
  namespace = "system"
}

output "alert_id" {
  value = data.f5_distributed_cloud_alert.example.id
}
