# Example configuration for f5_distributed_cloud_topology data source

data "f5_distributed_cloud_topology" "example" {
  name      = "existing-topology"
  namespace = "system"
}

output "topology_id" {
  value = data.f5_distributed_cloud_topology.example.id
}
