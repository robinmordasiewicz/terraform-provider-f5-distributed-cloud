# Example configuration for f5xc_user_token data source

data "f5xc_user_token" "example" {
  name      = "existing-user_token"
  namespace = "system"
}

output "user_token_id" {
  value = data.f5xc_user_token.example.id
}
