# Example configuration for f5xc_tpm_manager data source

data "f5xc_tpm_manager" "example" {
  name      = "existing-tpm_manager"
  namespace = "system"
}

output "tpm_manager_id" {
  value = data.f5xc_tpm_manager.example.id
}
