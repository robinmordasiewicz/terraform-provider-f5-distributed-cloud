# Example configuration for f5_distributed_cloud_crl data source

data "f5_distributed_cloud_crl" "example" {
  name      = "existing-crl"
  namespace = "system"
}

output "crl_id" {
  value = data.f5_distributed_cloud_crl.example.id
}
