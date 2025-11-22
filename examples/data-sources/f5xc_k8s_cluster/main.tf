# Example configuration for f5xc_k8s_cluster data source

data "f5xc_k8s_cluster" "example" {
  name      = "existing-k8s_cluster"
  namespace = "system"
}

output "k8s_cluster_id" {
  value = data.f5xc_k8s_cluster.example.id
}
