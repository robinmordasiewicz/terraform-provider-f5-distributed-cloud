# Example configuration for f5_distributed_cloud_usb data source

data "f5_distributed_cloud_usb" "example" {
  name      = "existing-usb"
  namespace = "system"
}

output "usb_id" {
  value = data.f5_distributed_cloud_usb.example.id
}
