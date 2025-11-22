resource "f5xc_k8s_cluster_role" "example" {
  name        = "my-cluster-role"
  namespace   = "system"
  description = "Cluster role for admin access"
}
