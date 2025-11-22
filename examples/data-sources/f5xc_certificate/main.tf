# Example configuration for f5xc_certificate data source

data "f5xc_certificate" "example" {
  name      = "existing-certificate"
  namespace = "system"
}

output "certificate_id" {
  value = data.f5xc_certificate.example.id
}
