# Example configuration for f5_distributed_cloud_advertise_policy

resource "f5_distributed_cloud_advertise_policy" "example" {
  name        = "example-advertise_policy"
  namespace   = "system"
  description = "Example AdvertisePolicy resource managed by Terraform"

  # Add additional configuration as needed
}
