# Example configuration for f5xc_usb_policy data source

data "f5xc_usb_policy" "example" {
  name      = "existing-usb_policy"
  namespace = "system"
}

output "usb_policy_id" {
  value = data.f5xc_usb_policy.example.id
}
