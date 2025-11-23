# Example configuration for f5_distributed_cloud_signup data source

data "f5_distributed_cloud_signup" "example" {
  name      = "existing-signup"
  namespace = "system"
}

output "signup_id" {
  value = data.f5_distributed_cloud_signup.example.id
}
