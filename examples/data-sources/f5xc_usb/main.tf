# Example configuration for f5xc_usb data source

data "f5xc_usb" "example" {
  name      = "existing-usb"
  namespace = "system"
}

output "usb_id" {
  value = data.f5xc_usb.example.id
}
