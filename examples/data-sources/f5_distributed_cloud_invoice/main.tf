# Example configuration for f5_distributed_cloud_invoice data source

data "f5_distributed_cloud_invoice" "example" {
  name      = "existing-invoice"
  namespace = "system"
}

output "invoice_id" {
  value = data.f5_distributed_cloud_invoice.example.id
}
