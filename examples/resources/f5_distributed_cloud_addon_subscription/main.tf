# Example configuration for f5_distributed_cloud_addon_subscription

resource "f5_distributed_cloud_addon_subscription" "example" {
  name        = "example-addon_subscription"
  namespace   = "system"
  description = "Example AddonSubscription resource managed by Terraform"

  # Add additional configuration as needed
}
