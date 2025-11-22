resource "f5xc_k8s_cluster_role_binding" "example" {
  name        = "my-cluster-role-binding"
  namespace   = "system"
  description = "Binding admin role to service account"
}
