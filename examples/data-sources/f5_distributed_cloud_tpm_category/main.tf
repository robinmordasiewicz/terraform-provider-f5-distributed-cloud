# Example configuration for f5_distributed_cloud_tpm_category data source

data "f5_distributed_cloud_tpm_category" "example" {
  name      = "existing-tpm_category"
  namespace = "system"
}

output "tpm_category_id" {
  value = data.f5_distributed_cloud_tpm_category.example.id
}
