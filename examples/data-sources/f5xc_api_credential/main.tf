# Example: Look up an existing API Credential
data "f5xc_api_credential" "example" {
  name      = "my-api-credential"
  namespace = "system"
}

output "credential_type" {
  value = data.f5xc_api_credential.example.credential_type
}
