# Example configuration for f5xc_signup data source

data "f5xc_signup" "example" {
  name      = "existing-signup"
  namespace = "system"
}

output "signup_id" {
  value = data.f5xc_signup.example.id
}
