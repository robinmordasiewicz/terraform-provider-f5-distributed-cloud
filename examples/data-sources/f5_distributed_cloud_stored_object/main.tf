# Example configuration for f5_distributed_cloud_stored_object data source

data "f5_distributed_cloud_stored_object" "example" {
  name      = "existing-stored_object"
  namespace = "system"
}

output "stored_object_id" {
  value = data.f5_distributed_cloud_stored_object.example.id
}
