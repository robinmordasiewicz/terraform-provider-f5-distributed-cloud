# Example configuration for f5xc_addon_subscription data source

data "f5xc_addon_subscription" "example" {
  name      = "existing-addon_subscription"
  namespace = "system"
}

output "addon_subscription_id" {
  value = data.f5xc_addon_subscription.example.id
}
