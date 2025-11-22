resource "f5xc_k8s_cluster" "example" {
  name        = "my-k8s-cluster"
  namespace   = "system"
  description = "Production Kubernetes cluster"
}
