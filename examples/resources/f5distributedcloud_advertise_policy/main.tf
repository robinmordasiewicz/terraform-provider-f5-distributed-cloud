# Example configuration for f5distributedcloud_advertise_policy

resource "f5distributedcloud_advertise_policy" "example" {
  name        = "example-advertise_policy"
  namespace   = "system"
  description = "Example AdvertisePolicy resource managed by Terraform"

  # Add additional configuration as needed
}
