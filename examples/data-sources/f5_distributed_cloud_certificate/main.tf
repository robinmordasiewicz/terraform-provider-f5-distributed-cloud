# Example configuration for f5_distributed_cloud_certificate data source

data "f5_distributed_cloud_certificate" "example" {
  name      = "existing-certificate"
  namespace = "system"
}

output "certificate_id" {
  value = data.f5_distributed_cloud_certificate.example.id
}
