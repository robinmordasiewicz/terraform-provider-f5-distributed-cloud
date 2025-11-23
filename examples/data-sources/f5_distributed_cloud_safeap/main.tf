# Example configuration for f5_distributed_cloud_safeap data source

data "f5_distributed_cloud_safeap" "example" {
  name      = "existing-safeap"
  namespace = "system"
}

output "safeap_id" {
  value = data.f5_distributed_cloud_safeap.example.id
}
