# Example configuration for f5xc_namespace data source

data "f5xc_namespace" "example" {
  name      = "existing-namespace"
  namespace = "system"
}

output "namespace_id" {
  value = data.f5xc_namespace.example.id
}
