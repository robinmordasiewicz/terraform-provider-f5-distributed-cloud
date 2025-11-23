# Example configuration for f5distributedcloud_namespace data source

data "f5distributedcloud_namespace" "example" {
  name      = "existing-namespace"
  namespace = "system"
}

output "namespace_id" {
  value = data.f5distributedcloud_namespace.example.id
}
