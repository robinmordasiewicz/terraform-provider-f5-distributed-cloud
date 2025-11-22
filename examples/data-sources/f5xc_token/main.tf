# Example configuration for f5xc_token data source

data "f5xc_token" "example" {
  name      = "existing-token"
  namespace = "system"
}

output "token_id" {
  value = data.f5xc_token.example.id
}
