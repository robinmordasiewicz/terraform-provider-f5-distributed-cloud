# Example configuration for f5_distributed_cloud_contact data source

data "f5_distributed_cloud_contact" "example" {
  name      = "existing-contact"
  namespace = "system"
}

output "contact_id" {
  value = data.f5_distributed_cloud_contact.example.id
}
