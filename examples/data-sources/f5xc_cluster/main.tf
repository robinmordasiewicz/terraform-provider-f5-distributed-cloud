# Example configuration for f5xc_cluster data source

data "f5xc_cluster" "example" {
  name      = "existing-cluster"
  namespace = "system"
}

output "cluster_id" {
  value = data.f5xc_cluster.example.id
}
