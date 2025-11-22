# Example configuration for f5xc_contact data source

data "f5xc_contact" "example" {
  name      = "existing-contact"
  namespace = "system"
}

output "contact_id" {
  value = data.f5xc_contact.example.id
}
