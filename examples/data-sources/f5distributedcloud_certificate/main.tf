# Example configuration for f5distributedcloud_certificate data source

data "f5distributedcloud_certificate" "example" {
  name      = "existing-certificate"
  namespace = "system"
}

output "certificate_id" {
  value = data.f5distributedcloud_certificate.example.id
}
