# Example configuration for f5xc_addon_subscription

resource "f5xc_addon_subscription" "example" {
  name        = "example-addon_subscription"
  namespace   = "system"
  description = "Example AddonSubscription resource managed by Terraform"

  # Add additional configuration as needed
}
