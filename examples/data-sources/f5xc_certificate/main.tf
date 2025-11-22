# Example: Look up an existing certificate
data "f5xc_certificate" "example" {
  name      = "my-certificate"
  namespace = "my-namespace"
}

# Output the certificate type (TLS or CA)
output "cert_type" {
  value = data.f5xc_certificate.example.cert_type
}

output "certificate_id" {
  value = data.f5xc_certificate.example.id
}
