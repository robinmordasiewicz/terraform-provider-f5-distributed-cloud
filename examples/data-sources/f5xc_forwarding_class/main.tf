# Example configuration for f5xc_forwarding_class data source

data "f5xc_forwarding_class" "example" {
  name      = "existing-forwarding_class"
  namespace = "system"
}

output "forwarding_class_id" {
  value = data.f5xc_forwarding_class.example.id
}
