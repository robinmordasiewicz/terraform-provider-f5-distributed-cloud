# Example configuration for f5xc_subscription data source

data "f5xc_subscription" "example" {
  name      = "existing-subscription"
  namespace = "system"
}

output "subscription_id" {
  value = data.f5xc_subscription.example.id
}
