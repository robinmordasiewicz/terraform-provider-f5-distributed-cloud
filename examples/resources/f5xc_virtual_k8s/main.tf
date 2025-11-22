resource "f5xc_virtual_k8s" "example" {
  name        = "my-vk8s"
  namespace   = "default"
  description = "Example virtual Kubernetes namespace"
}
