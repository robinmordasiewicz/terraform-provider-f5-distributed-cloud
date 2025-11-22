# Example configuration for f5xc_virtual_k8s

resource "f5xc_virtual_k8s" "example" {
  name        = "example-virtual_k8s"
  namespace   = "system"
  description = "Example VirtualK8S resource managed by Terraform"

  # Add additional configuration as needed
}
