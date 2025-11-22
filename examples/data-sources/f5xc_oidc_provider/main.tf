# Example configuration for f5xc_oidc_provider data source

data "f5xc_oidc_provider" "example" {
  name      = "existing-oidc_provider"
  namespace = "system"
}

output "oidc_provider_id" {
  value = data.f5xc_oidc_provider.example.id
}
