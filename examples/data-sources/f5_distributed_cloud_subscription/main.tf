# Example configuration for f5_distributed_cloud_subscription data source

data "f5_distributed_cloud_subscription" "example" {
  name      = "existing-subscription"
  namespace = "system"
}

output "subscription_id" {
  value = data.f5_distributed_cloud_subscription.example.id
}
