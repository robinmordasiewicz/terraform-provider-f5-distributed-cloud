# Example configuration for f5xc_stored_object data source

data "f5xc_stored_object" "example" {
  name      = "existing-stored_object"
  namespace = "system"
}

output "stored_object_id" {
  value = data.f5xc_stored_object.example.id
}
