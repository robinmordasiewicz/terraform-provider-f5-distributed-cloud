# Example configuration for f5xc_api_credential data source

data "f5xc_api_credential" "example" {
  name      = "existing-api_credential"
  namespace = "system"
}

output "api_credential_id" {
  value = data.f5xc_api_credential.example.id
}
