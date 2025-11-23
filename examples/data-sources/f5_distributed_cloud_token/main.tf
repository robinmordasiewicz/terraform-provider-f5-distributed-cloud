# Example configuration for f5_distributed_cloud_token data source

data "f5_distributed_cloud_token" "example" {
  name      = "existing-token"
  namespace = "system"
}

output "token_id" {
  value = data.f5_distributed_cloud_token.example.id
}
