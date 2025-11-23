# Example configuration for f5distributedcloud_token data source

data "f5distributedcloud_token" "example" {
  name      = "existing-token"
  namespace = "system"
}

output "token_id" {
  value = data.f5distributedcloud_token.example.id
}
