# Example configuration for f5xc_virtual_k8s data source

data "f5xc_virtual_k8s" "example" {
  name      = "existing-virtual_k8s"
  namespace = "system"
}

output "virtual_k8s_id" {
  value = data.f5xc_virtual_k8s.example.id
}
