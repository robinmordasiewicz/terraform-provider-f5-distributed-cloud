# Example configuration for f5xc_tpm_api_key data source

data "f5xc_tpm_api_key" "example" {
  name      = "existing-tpm_api_key"
  namespace = "system"
}

output "tpm_api_key_id" {
  value = data.f5xc_tpm_api_key.example.id
}
