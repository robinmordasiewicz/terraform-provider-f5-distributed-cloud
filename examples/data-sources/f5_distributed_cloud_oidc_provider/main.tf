# Example configuration for f5_distributed_cloud_oidc_provider data source

data "f5_distributed_cloud_oidc_provider" "example" {
  name      = "existing-oidc_provider"
  namespace = "system"
}

output "oidc_provider_id" {
  value = data.f5_distributed_cloud_oidc_provider.example.id
}
