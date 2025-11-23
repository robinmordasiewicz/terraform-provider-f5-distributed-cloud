# Example configuration for f5_distributed_cloud_customer_support data source

data "f5_distributed_cloud_customer_support" "example" {
  name      = "existing-customer_support"
  namespace = "system"
}

output "customer_support_id" {
  value = data.f5_distributed_cloud_customer_support.example.id
}
