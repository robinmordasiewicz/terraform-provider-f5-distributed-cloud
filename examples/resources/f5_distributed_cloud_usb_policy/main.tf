# Example configuration for f5_distributed_cloud_usb_policy

resource "f5_distributed_cloud_usb_policy" "example" {
  name        = "example-usb_policy"
  namespace   = "system"
  description = "Example UsbPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
