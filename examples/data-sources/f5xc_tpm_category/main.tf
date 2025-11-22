# Example configuration for f5xc_tpm_category data source

data "f5xc_tpm_category" "example" {
  name      = "existing-tpm_category"
  namespace = "system"
}

output "tpm_category_id" {
  value = data.f5xc_tpm_category.example.id
}
