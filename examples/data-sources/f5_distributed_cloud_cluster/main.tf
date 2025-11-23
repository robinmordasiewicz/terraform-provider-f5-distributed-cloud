# Example configuration for f5_distributed_cloud_cluster data source

data "f5_distributed_cloud_cluster" "example" {
  name      = "existing-cluster"
  namespace = "system"
}

output "cluster_id" {
  value = data.f5_distributed_cloud_cluster.example.id
}
