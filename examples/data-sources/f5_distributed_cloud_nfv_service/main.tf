# Example configuration for f5_distributed_cloud_nfv_service data source

data "f5_distributed_cloud_nfv_service" "example" {
  name      = "existing-nfv_service"
  namespace = "system"
}

output "nfv_service_id" {
  value = data.f5_distributed_cloud_nfv_service.example.id
}
