# Example configuration for f5xc_customer_support data source

data "f5xc_customer_support" "example" {
  name      = "existing-customer_support"
  namespace = "system"
}

output "customer_support_id" {
  value = data.f5xc_customer_support.example.id
}
