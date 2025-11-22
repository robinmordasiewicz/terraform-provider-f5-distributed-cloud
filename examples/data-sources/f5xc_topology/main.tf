# Example configuration for f5xc_topology data source

data "f5xc_topology" "example" {
  name      = "existing-topology"
  namespace = "system"
}

output "topology_id" {
  value = data.f5xc_topology.example.id
}
