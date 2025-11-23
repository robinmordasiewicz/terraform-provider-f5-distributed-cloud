# Example configuration for f5_distributed_cloud_forwarding_class data source

data "f5_distributed_cloud_forwarding_class" "example" {
  name      = "existing-forwarding_class"
  namespace = "system"
}

output "forwarding_class_id" {
  value = data.f5_distributed_cloud_forwarding_class.example.id
}
