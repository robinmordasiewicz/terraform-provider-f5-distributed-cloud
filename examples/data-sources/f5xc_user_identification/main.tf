# Example configuration for f5xc_user_identification data source

data "f5xc_user_identification" "example" {
  name      = "existing-user_identification"
  namespace = "system"
}

output "user_identification_id" {
  value = data.f5xc_user_identification.example.id
}
