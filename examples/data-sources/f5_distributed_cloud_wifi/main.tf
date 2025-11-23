# Example configuration for f5_distributed_cloud_wifi data source

data "f5_distributed_cloud_wifi" "example" {
  name      = "existing-wifi"
  namespace = "system"
}

output "wifi_id" {
  value = data.f5_distributed_cloud_wifi.example.id
}
