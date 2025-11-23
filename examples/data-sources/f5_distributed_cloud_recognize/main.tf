# Example configuration for f5_distributed_cloud_recognize data source

data "f5_distributed_cloud_recognize" "example" {
  name      = "existing-recognize"
  namespace = "system"
}

output "recognize_id" {
  value = data.f5_distributed_cloud_recognize.example.id
}
