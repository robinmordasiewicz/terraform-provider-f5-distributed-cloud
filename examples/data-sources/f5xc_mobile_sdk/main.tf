# Example configuration for f5xc_mobile_sdk data source

data "f5xc_mobile_sdk" "example" {
  name      = "existing-mobile_sdk"
  namespace = "system"
}

output "mobile_sdk_id" {
  value = data.f5xc_mobile_sdk.example.id
}
