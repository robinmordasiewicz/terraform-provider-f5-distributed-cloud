# Example configuration for f5xc_safeap data source

data "f5xc_safeap" "example" {
  name      = "existing-safeap"
  namespace = "system"
}

output "safeap_id" {
  value = data.f5xc_safeap.example.id
}
