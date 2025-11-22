# Example configuration for f5xc_invoice data source

data "f5xc_invoice" "example" {
  name      = "existing-invoice"
  namespace = "system"
}

output "invoice_id" {
  value = data.f5xc_invoice.example.id
}
